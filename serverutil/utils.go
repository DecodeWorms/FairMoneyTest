package serverutil

import (
	"context"
	server "fairmoneytest/Server"
	"fairmoneytest/config"
	"fairmoneytest/handler"
	"fairmoneytest/storage"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// SetUpDatabase sets up database with dependencies
func SetUpDatabase(url, name string) (storage.DataStore, *mongo.Client) {
	repo, client, err := storage.New(url, name)
	if err != nil {
		log.Fatalf("Error failed to open MongoDB: %v", err)
	}
	return repo, client
}

// SetUpHandler sets up handler with dependencies
func SetUpHandler(store storage.DataStore) handler.TransactionHandler {
	return handler.NewTransactionHandler(store)
}

// SetUpServer sets up server with dependencies
func SetUpServer(userHandler handler.TransactionHandler) server.TransactionServer {
	return server.NewTransactionServer(userHandler)
}

// SetupRouter sets up router with dependencies
func SetupRouter(server *server.TransactionServer) *gin.Engine {
	router := gin.Default()

	// Add Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	//Api endpoints are here
	router.POST("/transaction/credit", server.Credit())
	router.POST("/transaction/debit", server.Debit())
	router.GET("/transaction/transaction/:reference", server.GetTransaction())

	return router
}

// StartServer starts the backend services
func StartServer(router *gin.Engine, client *mongo.Client) {
	var c config.Config
	c = config.ImportConfig(config.OSSource{})
	interruptHandler := make(chan os.Signal, 1)
	signal.Notify(interruptHandler, syscall.SIGTERM, syscall.SIGINT)

	addr := fmt.Sprintf(":%s", c.ServicePort)
	go func(addr string) {
		log.Println(fmt.Sprintf("Jacq API service running on %v. Environment=%s", addr, c.AppEnv))
		if err := http.ListenAndServe(addr, router); err != nil {
			log.Printf("Error starting server: %v", err)
		}
	}(addr)

	<-interruptHandler
	log.Println("Closing application...")
	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatal("Failed to disconnect from database")
	}
}

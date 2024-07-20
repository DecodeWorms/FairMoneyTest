package server

import (
	"bytes"
	"encoding/json"
	"fairmoneytest/handler"
	"fairmoneytest/model"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type TransactionServer struct {
	handler handler.TransactionHandler
}

func NewTransactionServer(handler handler.TransactionHandler) TransactionServer {
	return TransactionServer{handler: handler}
}

// Credit Handles credit request
func (server TransactionServer) Credit() gin.HandlerFunc {
	return func(context *gin.Context) {
		//Third-party api endpoint
		url := "http://third-party/payments"

		var trx model.Transactions

		if err := context.ShouldBindJSON(&trx); err != nil {
			context.JSON(http.StatusBadRequest, "Invalid request body")
			return
		}

		//Ensure that a valid amount is received
		if trx.Amount <= 0 {
			context.JSON(http.StatusBadRequest, "Invalid amount")
			return
		}

		//Convert the request body to json
		jsData, err := json.Marshal(trx)
		if err != nil {
			context.JSON(http.StatusBadRequest, "Unable to Marshal the request body")
			return
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsData))
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initiate a request response from third-party API"})
			return
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute a request from third-party API"})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			context.JSON(http.StatusInternalServerError, "Internal server error")
			return
		}

		//If everything goes well, keep record of the transaction
		rec, err := server.handler.RecordCreditTransaction(&model.Transactions{
			AccountID: trx.AccountID,
			Reference: trx.Reference,
			Amount:    trx.Amount,
		})

		/*if err != nil {
			context.JSON(http.StatusInternalServerError, "Internal server error, unable to record the tra")
			return
		}
		*
		*/
		context.JSON(http.StatusOK, gin.H{"data": rec, "success": "successfull"})
	}
}

// Debit handles debit request
func (server TransactionServer) Debit() gin.HandlerFunc {
	return func(context *gin.Context) {
		//Third-party api endpoint
		url := "http://third-party/payments"

		var trx model.Transactions

		if err := context.ShouldBindJSON(&trx); err != nil {
			context.JSON(http.StatusBadRequest, "Invalid request body")
			return
		}

		//Ensure that a valid amount is received .
		if trx.Amount <= 0 {
			context.JSON(http.StatusBadRequest, "Invalid amount")
			return
		}

		//Convert the request body to json
		jsData, err := json.Marshal(trx)
		if err != nil {
			context.JSON(http.StatusBadRequest, "Unable to Marshal the request body")
			return
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsData))
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initiate a request from third-party API"})
			return
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute a request from third-party API"})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			context.JSON(http.StatusInternalServerError, "Internal server error")
			return
		}

		//If everything goes well, keep record of the transaction
		rec, err := server.handler.RecordDebitTransaction(&model.Transactions{
			AccountID: trx.AccountID,
			Reference: trx.Reference,
			Amount:    trx.Amount,
		})

		if err != nil {
			context.JSON(http.StatusInternalServerError, "Internal server error")
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": rec})
	}
}

func (server TransactionServer) GetTransaction() gin.HandlerFunc {
	return func(context *gin.Context) {
		//Extract transaction reference from the request
		ref := context.Param("reference")

		//Third-party payment
		url := "http://third-party/payments/:" + ref

		resp, err := http.Get(url)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get response from third-party API"})
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
			return
		}

		var transaction model.Transactions
		err = json.Unmarshal(body, &transaction)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response body"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"transaction": transaction})
	}
}

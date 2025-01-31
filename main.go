package main

import (
	"fairmoneytest/config"
	"fairmoneytest/serverutil"
)

var c config.Config

func main() {
	c = config.ImportConfig(config.OSSource{})
	store, client := serverutil.SetUpDatabase(c.DatabaseURL, c.DatabaseName)
	handler := serverutil.SetUpHandler(store)
	server := serverutil.SetUpServer(handler)
	router := serverutil.SetupRouter(&server)
	serverutil.StartServer(router, client)
}

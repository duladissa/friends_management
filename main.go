package main

import (
	"net/http"

	"github.com/duladissa/friends_management/handler"
	"github.com/kataras/golog"
	"github.com/spf13/viper"
)

func main() {
	c := viper.New()
	c.SetDefault("PORT", "8888")
	c.SetDefault("LOG_LEVEL", "debug")
	c.SetDefault("MONGODB_URI", "mongodb://localhost:27017")
	c.SetDefault("MONGODB_NAME", "friends_management")

	//Read Env variables
	c.AutomaticEnv()

	serverAddress := c.Get("PORT").(string)
	logLevel := c.Get("LOG_LEVEL").(string)
	golog.SetLevel(logLevel)

	server := http.Server{Addr: ":" + serverAddress, Handler: handler.InitHandlersAndRoutes(c)}

	golog.Info("Server started at", ":"+serverAddress)

	err := server.ListenAndServe()

	if err != nil {
		golog.Fatal(err)
	}
}

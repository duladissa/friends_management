package main

import (
	"net/http"
	"github.com/kataras/golog"
	"github.com/spf13/viper"
	"github.com/duladissa/friends_management/handler"
)

func main() {
	c := viper.New()
	c.SetDefault("PORT", "8888")
	c.SetDefault("LOG_LEVEL", "debug")

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

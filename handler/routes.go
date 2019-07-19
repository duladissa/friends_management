package handler

import (
	"net/http"

	"github.com/duladissa/friends_management/restapi"
	"github.com/kataras/golog"
	"github.com/spf13/viper"
)

//InitHandlersAndRoutes ... All the routes
func InitHandlersAndRoutes(envConfig *viper.Viper) http.Handler {
	config := restapi.Config{
		FriendsAPI: NewFriendsAPI(),
		Logger:     golog.Debugf,
	}
	handler, err := restapi.Handler(config)

	if err != nil {
		golog.Fatal(err)
	}
	return handler
}

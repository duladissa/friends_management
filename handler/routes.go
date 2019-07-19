package handler

import (
	"net/http"

	"github.com/duladissa/friends_management/database"
	"github.com/duladissa/friends_management/restapi"
	"github.com/kataras/golog"
	"github.com/spf13/viper"
)

//InitHandlersAndRoutes ... All the routes
func InitHandlersAndRoutes(envConfig *viper.Viper) http.Handler {
	dbURI := envConfig.Get("MONGODB_URI").(string)
	dbName := envConfig.Get("MONGODB_NAME").(string)

	database, _ := database.CreateDatabase(dbURI, dbName)
	config := restapi.Config{
		FriendsAPI: NewFriendsAPI(database),
		Logger:     golog.Debugf,
	}
	handler, err := restapi.Handler(config)

	if err != nil {
		golog.Fatal(err)
	}
	return handler
}

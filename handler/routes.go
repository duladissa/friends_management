package handler

import(
	"net/http"
	"github.com/spf13/viper"
	"github.com/kataras/golog"
	"github.com/duladissa/friends_management/restapi"
)

//InitHandlersAndRoutes ... All the routes
func InitHandlersAndRoutes(envConfig *viper.Viper) http.Handler {
	config := restapi.Config{
		Logger:          golog.Debugf,
	}
	handler, err := restapi.Handler(config)

	if err != nil {
		golog.Fatal(err)
	}
	return handler
}
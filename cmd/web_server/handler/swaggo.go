package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggo_gen "github.com/oj-lab/platform/cmd/web_server/swaggo_gen"
	"github.com/spf13/viper"
	swagger_files "github.com/swaggo/files"
	gin_swagger "github.com/swaggo/gin-swagger"
)

const (
	servicePortProp = "service.port"
	serviceHostProp = "service.host"
)

var (
	swaggerHost string
)

func SetupSwaggoRouter(r *gin.RouterGroup) {
	r.GET("/swagger/*any", gin_swagger.WrapHandler(swagger_files.Handler))
}

func init() {
	sevicePort := viper.GetUint(servicePortProp)
	seviceHost := viper.GetString(serviceHostProp)
	swaggerHost = fmt.Sprintf("%s:%d", seviceHost, sevicePort)
	println("Swagger host is set to: " + swaggerHost)
	// programmatically set swagger info
	swaggo_gen.SwaggerInfo.Title = "OJ Lab Services API"
	swaggo_gen.SwaggerInfo.Version = "1.0"
	swaggo_gen.SwaggerInfo.Host = swaggerHost
	swaggo_gen.SwaggerInfo.BasePath = "/api/v1"
	swaggo_gen.SwaggerInfo.Schemes = []string{"http"}
}

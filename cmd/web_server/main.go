package main

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/oj-lab/oj-lab-platform/cmd/web_server/handler"

	"github.com/gin-gonic/gin"
	"github.com/oj-lab/oj-lab-platform/modules/config"
	"github.com/oj-lab/oj-lab-platform/modules/log"
	"github.com/oj-lab/oj-lab-platform/modules/middleware"
)

const (
	serviceForceConsoleColorProp = "service.force_console_color"
	servicePortProp              = "service.port"
	serviceModeProp              = "service.mode"
	swaggerOnProp                = "service.swagger_on"
	frontendDistProp             = "service.frontend_dist"
)

var (
	serviceForceConsoleColor bool
	servicePort              string
	serviceMode              string
	swaggerOn                bool
	frontendDist             string
)

func init() {
	serviceForceConsoleColor = config.AppConfig.GetBool(serviceForceConsoleColorProp)
	servicePort = config.AppConfig.GetString(servicePortProp)
	serviceMode = config.AppConfig.GetString(serviceModeProp)
	swaggerOn = config.AppConfig.GetBool(swaggerOnProp)
	frontendDist = config.AppConfig.GetString(frontendDistProp)
}

func GetProjectDir() string {
	_, b, _, _ := runtime.Caller(0)
	projectDir := filepath.Join(filepath.Dir(b), "..", "..")

	return projectDir
}

func main() {
	if serviceForceConsoleColor {
		gin.ForceConsoleColor()
	}
	r := gin.Default()
	r.Use(middleware.HandleError)
	gin.SetMode(serviceMode)

	baseRouter := r.Group("/")
	if frontendDist != "" {
		// If dist folder is not empty, serve frontend
		if _, err := os.Stat(frontendDist); os.IsNotExist(err) {
			log.AppLogger().Warn("Frontend dist is set but folder not found")
		} else {
			log.AppLogger().Info("Serving frontend...")
			r.LoadHTMLFiles(frontendDist + "/index.html")
			handler.SetupFrontendRoute(baseRouter, frontendDist)
		}
	}

	if swaggerOn {
		log.AppLogger().Info("Serving swagger Doc...")
		handler.SetupSwaggoRouter(baseRouter)
	}

	apiRouter := r.Group("/api/v1")
	handler.SetupUserRouter(apiRouter)
	handler.SetupProblemRouter(apiRouter)
	handler.SetupEventRouter(apiRouter)
	handler.SetupJudgeRouter(apiRouter)
	handler.SetupJudgeTaskRouter(apiRouter)
	handler.SetupJudgeResultRouter(apiRouter)

	err := r.Run(servicePort)
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

package router

import (
	"io"
	"os"
	"router-gostashlg-template/delivery/http/router/middleware"
	"router-gostashlg-template/entities/common/logger"

	"github.com/gin-gonic/gin"
)

func Start() error {
	gin.SetMode(gin.ReleaseMode)

	//Discard semua output yang dicatat oleh gin karena print out akan dicetak sesuai kebutuhan programmer
	gin.DefaultWriter = io.Discard

	router := gin.Default() //create router engine by default
	router.Use(gin.Recovery(), middleware.RequestLogger, middleware.ResponseLogger)

	RegisterHandler(router)
	listenerPort := os.Getenv("app.listener_port")
	logger.PrintLogf("[HTTP] Listening at : %s", listenerPort)
	return router.Run(":" + listenerPort)

}

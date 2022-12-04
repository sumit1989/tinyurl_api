package router

import (
	conf "tinyurl_api/configuration"
	"tinyurl_api/controller"
	"tinyurl_api/service"

	"github.com/gin-gonic/gin"

	logger "github.com/sirupsen/logrus"
)

func Setup(conf conf.Config) *gin.Engine {
	logger.Info("In Setup Function.")
	router := gin.Default()
	service := service.NewShortenUrlService()
	controller := controller.NewShortenUrlController(service)
	internalAuthRoute := router.Group("/internal/test/v1/")
	internalAuthRoute.GET("/url", controller.CreateShortenUrl)
	return router

}

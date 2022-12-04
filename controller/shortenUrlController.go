package controller

import (
	"net/http"
	model "tinyurl_api/common"
	service "tinyurl_api/service"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

type shortenUrlController struct {
	service service.ShortenUrlService
}

func NewShortenUrlController(service service.ShortenUrlService) shortenUrlController {
	return shortenUrlController{service: service}

}

func (s *shortenUrlController) CreateShortenUrl(ctx *gin.Context) {
	logger.Info("In CreateShortenUrl Method Controller.")
	response, err := s.service.CreateShortenUrl(ctx)
	if (err != model.Error{}) {
		ctx.AbortWithStatusJSON(err.ErrorCode, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

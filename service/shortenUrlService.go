package service

import (
	"net/url"
	model "tinyurl_api/common"
	"tinyurl_api/utility"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

type shortenUrlService struct {
}

type ShortenUrlService interface {
	CreateShortenUrl(ctx *gin.Context) (model.Response, model.Error)
}

func NewShortenUrlService() ShortenUrlService {
	return &shortenUrlService{}
}

func (s *shortenUrlService) CreateShortenUrl(ctx *gin.Context) (model.Response, model.Error) {
	logger.Info("In CreateShortenUrl Method.")
	var response string
	requestParameter := ctx.Request.URL.Query().Get("longUrl")
	if requestParameter == "" {
		return model.Response{}, model.Error{ErrorCode: 400, ErrorMssage: "Required_Parameter_Missing"}
	}
	_, errParseUri := url.ParseRequestURI(requestParameter)

	if errParseUri != nil {
		return model.Response{}, model.Error{ErrorCode: 400, ErrorMssage: "Inappropriate_URL_Format"}
	}
	response, err := utility.GenerateHashAndInsert(requestParameter, 5)
	if err != nil {
		return model.Response{}, model.Error{ErrorCode: 500, ErrorMssage: "Internal_Server_Error"}
	}
	logger.Info("Successfully created the short url.")
	return model.Response{ResponseCode: "200", ResponseMssage: utility.BaseUrl + response}, model.Error{}

}

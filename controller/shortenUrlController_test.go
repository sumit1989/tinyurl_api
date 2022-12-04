package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
	model "tinyurl_api/common"
	"tinyurl_api/mocks"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
)

type shortenUrlControllerTestSuite struct {
	suite.Suite
	mockCtrl             *gomock.Controller
	recorder             *httptest.ResponseRecorder
	ginContext           *gin.Context
	shortenService       *mocks.MockShortenUrlService
	shortenUrlController shortenUrlController
	responseRecorder     *httptest.ResponseRecorder
}

func TestAuditMonitorTestSuiteTestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	suite.Run(t, new(shortenUrlControllerTestSuite))
}

func (suite *shortenUrlControllerTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.recorder = httptest.NewRecorder()
	suite.ginContext, _ = gin.CreateTestContext(suite.recorder)
	suite.ginContext.Request, _ = http.NewRequest("GET", "", nil)
	suite.shortenService = mocks.NewMockShortenUrlService(suite.mockCtrl)
	suite.shortenUrlController = NewShortenUrlController(suite.shortenService)
	suite.responseRecorder = httptest.NewRecorder()
	suite.ginContext, _ = gin.CreateTestContext(suite.responseRecorder)

}

func (suite shortenUrlControllerTestSuite) TestCreateShortenUrlSuccessCase() {
	response := model.Response{
		ResponseCode:   "200",
		ResponseMssage: "abcd",
	}
	errResponse := model.Error{}

	apiUrl := "/internal/test/v1/url?longUrl=http//yahoo.com/search/abc"
	suite.ginContext.Request, _ = http.NewRequest(http.MethodGet, apiUrl, nil)
	suite.shortenService.EXPECT().CreateShortenUrl(suite.ginContext).Return(response, errResponse)
	suite.shortenUrlController.CreateShortenUrl(suite.ginContext)
	Expect(suite.responseRecorder.Code).To(Equal(http.StatusOK))
}

func (suite shortenUrlControllerTestSuite) TestCreateShortenUrlErrorCase() {
	response := model.Response{}
	errResponse := model.Error{ErrorCode: 400, ErrorMssage: "Inappropriate_URL_Format"}

	apiUrl := "/internal/test/v1/url?longUrl=yahoo.com/search/abc"
	suite.ginContext.Request, _ = http.NewRequest(http.MethodGet, apiUrl, nil)
	suite.shortenService.EXPECT().CreateShortenUrl(suite.ginContext).Return(response, errResponse)
	suite.shortenUrlController.CreateShortenUrl(suite.ginContext)
	Expect(suite.responseRecorder.Code).To(Equal(http.StatusBadRequest))
}

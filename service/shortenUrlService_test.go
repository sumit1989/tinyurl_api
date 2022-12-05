package service

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
)

type shortenUrlServiceTestSuite struct {
	suite.Suite
	mockCtrl          *gomock.Controller
	recorder          *httptest.ResponseRecorder
	ginContext        *gin.Context
	shortenUrlService ShortenUrlService
	testContext       context.Context
}

func TestAuditMonitorTestSuiteTestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	suite.Run(t, new(shortenUrlServiceTestSuite))
}

func (suite *shortenUrlServiceTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.recorder = httptest.NewRecorder()
	suite.ginContext, _ = gin.CreateTestContext(suite.recorder)
	suite.ginContext.Request, _ = http.NewRequest("GET", "", nil)
	suite.shortenUrlService = NewShortenUrlService()
	suite.testContext = context.Background()
	//NewMockShortenUrlService(suite.mockCtrl)
	//suite.shortenUrlController = NewShortenUrlController(suite.shortenService)
	//suite.responseRecorder = httptest.NewRecorder()
	//suite.ginContext, _ = gin.CreateTestContext(suite.responseRecorder)

}

func (suite shortenUrlServiceTestSuite) TestCreateShortenUrlSuccessCase() {
	apiUrl := "/internal/test/v1/url?longUrl=http://yahoo.com/search/abc"
	suite.ginContext.Request, _ = http.NewRequest(http.MethodGet, apiUrl, nil)
	response, err := suite.shortenUrlService.CreateShortenUrl(suite.ginContext)
	suite.Empty(err.ErrorMssage)
	Expect(response.ResponseCode).To(Equal("200"))
	Expect(response.ResponseMssage).To(Equal("http://tinyurl/zM2NDE"))

}

func (suite shortenUrlServiceTestSuite) TestCreateShortenUrlErrorCaseRequiredParameterMissing() {
	apiUrl := "/internal/test/v1/url?longUrls=http://yahoo.com/search/abc"
	suite.ginContext.Request, _ = http.NewRequest(http.MethodGet, apiUrl, nil)
	response, err := suite.shortenUrlService.CreateShortenUrl(suite.ginContext)
	suite.Empty(response.ResponseMssage)
	Expect(err.ErrorCode).To(Equal(400))
	Expect(err.ErrorMssage).To(Equal("Required_Parameter_Missing"))

}

func (suite shortenUrlServiceTestSuite) TestCreateShortenUrlErrorCaseRequiredWrongUrlFormat() {
	apiUrl := "/internal/test/v1/url?longUrl=http//yahoo.com/search/abc"
	suite.ginContext.Request, _ = http.NewRequest(http.MethodGet, apiUrl, nil)
	response, err := suite.shortenUrlService.CreateShortenUrl(suite.ginContext)
	suite.Empty(response.ResponseMssage)
	Expect(err.ErrorCode).To(Equal(400))
	Expect(err.ErrorMssage).To(Equal("Inappropriate_URL_Format"))

}

package utility

import (
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
)

type utilTestSuite struct {
	suite.Suite
	ginContext *gin.Context
	mockCtrl   *gomock.Controller
	recorder   *httptest.ResponseRecorder
}

func TestUtilTestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	suite.Run(t, new(utilTestSuite))
}

func (suite utilTestSuite) TestLoadFileNotExist() {

	expected := "open test: no such file or directory"
	actual := LoadMap("test")
	suite.Equal(expected, actual.Error())
}

func (suite utilTestSuite) TestLoadFileExist() {
	actual := LoadMap(testFilePath)
	suite.Equal(nil, actual)
}

func (suite utilTestSuite) TestLoadMapData() {
	actual := LoadMap(testFilePath)
	suite.Equal(2, len(UrlMapping))
	for key, element := range UrlMapping {
		if key == "http://yahoo.com/search/abc" {
			suite.Equal("zM2NDE", element)
		} else if key == "https://yahoo.com/search/abc" {
			suite.Equal("2M2ZDM", element)
		}

	}
	suite.Equal(nil, actual)
}

func (suite utilTestSuite) TestGenerateHashAndErrorCase() {
	suite.recorder = httptest.NewRecorder()
	suite.ginContext, _ = gin.CreateTestContext(suite.recorder)
	_, err := GenerateHashAndInsert(suite.ginContext, "test", 5000, RegexExp)
	Expect(err.Error()).To(Equal("Unable to generate tiny URL"))

}

func (suite utilTestSuite) TestGenerateHashAndErrorCaseRegularExpression() {
	RegexExp := `^.*(?=.{7,})(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[^a-zA-Z0-9]).*$`
	suite.recorder = httptest.NewRecorder()
	suite.ginContext, _ = gin.CreateTestContext(suite.recorder)
	_, err := GenerateHashAndInsert(suite.ginContext, "http://yahoo.com/search/abc", 5, RegexExp)
	Expect(err.Error()).To(Equal("Unable to generate regular expression"))

}

func (suite utilTestSuite) TestGenerateHashAndSuccessCase() {
	var urlMapTest = make(map[string]string)
	suite.recorder = httptest.NewRecorder()
	suite.ginContext, _ = gin.CreateTestContext(suite.recorder)
	actual, err := GenerateHashAndInsert(suite.ginContext, "http://yahoo.com/search/abc", 5, RegexExp)
	suite.Equal(nil, err)
	suite.Equal("zM2NDE", actual)
	urlMapTest["http://yahoo.com/search/abc"] = "zM2NDE"
	eq := reflect.DeepEqual(UrlMapping, urlMapTest)

	suite.True(eq)

}

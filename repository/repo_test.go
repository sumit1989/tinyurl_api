package repository

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
)

type repositoryTestSuite struct {
	suite.Suite
	ginContext *gin.Context
	mockCtrl   *gomock.Controller
	recorder   *httptest.ResponseRecorder
}

func TestUtilTestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	suite.Run(t, new(repositoryTestSuite))
}

func (suite repositoryTestSuite) TestWriteFileNotExist() {

	expected := "open test: no such file or directory"
	err := WriteFile("test", "http://yahoo.com/search/abc", "zM2NDE")
	suite.Equal(expected, err.Error())
}

func (suite repositoryTestSuite) TestWriteSuccess() {
	err := WriteFile("repo.txt", "http://yahoo.com/search/abc", "zM2NDE")
	suite.Equal(nil, err)
}

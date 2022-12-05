//go:build integration
// +build integration

package integrationtest

import (
	"bytes"
	"fmt"

	"net/http"
	"net/http/httptest"

	"testing"
	conf "tinyurl_api/configuration"
	router "tinyurl_api/router"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
)

type tinyUrlIntegrationTestSuite struct {
	suite.Suite
	server     *httptest.Server
	configData conf.Config
}

func TestUtilTestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	suite.Run(t, new(tinyUrlIntegrationTestSuite))
}

func (suite tinyUrlIntegrationTestSuite) SetupSuite() {
	var err error
	suite.server, err = getMockServer("127.0.0.1:8083")
	if err != nil {
		suite.T().Fail()
	}

	suite.server.Start()
}

func (suite tinyUrlIntegrationTestSuite) TestTinyUrlSuccess() {
	conf, err := conf.Configuration("../configuration/config.json")
	suite.Nil(err)

	r := router.Setup(conf)

	w, err := performRequest(r, "GET", "/internal/test/v1/url?longUrl=http://yahoo.com/dhddh1", bytes.NewBuffer(successBody))
	if err != nil {
		suite.T().Fail()
	}
	suite.Equal(http.StatusOK, w.Code)
	suite.Equal("{\"responseCode\":\"200\",\"responseMssage\":\"http://tinyurl/DQzMWE\"}", w.Body.String())
}

func (suite tinyUrlIntegrationTestSuite) TearDownSuite() {
	if suite.server != nil {
		suite.server.Close()
	}
}

func (suite tinyUrlIntegrationTestSuite) TestTinyUrlFailure() {
	conf, err := conf.Configuration("../configuration/config.json")
	suite.Nil(err)
	r := router.Setup(conf)

	w, err := performRequest(r, "GET", "/internal/test/v1/url?longUrls=http://yahoo.com/dhddh1", bytes.NewBuffer(successBody))
	if err != nil {
		suite.T().Fail()
	}
	fmt.Println(w.Code)
	suite.Equal(http.StatusBadRequest, w.Code)
	suite.Equal("{\"errorCode\":400,\"errorMssage\":\"Required_Parameter_Missing\"}", w.Body.String())
}

var successBody = []byte(`"responseCode":"200","responseMssage":"http://tinyurl/DQzMWE"}`)

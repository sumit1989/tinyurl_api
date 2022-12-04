package utility

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MaskUtilTestSuite struct {
	suite.Suite
}

func TestUtilTestSuite(t *testing.T) {
	suite.Run(t, new(MaskUtilTestSuite))
}

func (suite MaskUtilTestSuite) TestLoadFileNotExist() {

	expected := "open test: no such file or directory"
	actual := LoadMap("test")
	fmt.Println(actual)
	suite.Equal(expected, actual.Error())
}

func (suite MaskUtilTestSuite) TestLoadFileExist() {

	actual := LoadMap(testFilePath)
	fmt.Println(actual)

	suite.Equal(nil, actual)
}

func (suite MaskUtilTestSuite) TestLoadMapData() {
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

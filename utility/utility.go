package utility

import (
	"bufio"
	"context"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"tinyurl_api/repository"

	logger "github.com/sirupsen/logrus"
)

var UrlMapping = make(map[string]string)

func GenerateHashAndInsert(ctxt context.Context, longURL string, startIndex int, regexExpr string) (string, error) {
	logger.Info("In GenerateHashAndInsert Function,Url is: ", longURL)
	byteURLData := []byte(longURL)
	hashedURLData := fmt.Sprintf("%x", md5.Sum(byteURLData))
	tinyURLRegex, err := regexp.Compile(regexExpr)
	if err != nil {
		logger.Error("Failed in compiling the regex", tinyURLRegex)
		return " ", errors.New("Unable to generate regular expression")
	}
	tinyURLData := tinyURLRegex.ReplaceAllString(base64.URLEncoding.EncodeToString([]byte(hashedURLData)), "_")
	logger.Info("aa", len(tinyURLData))
	if len(tinyURLData) < (startIndex + 6) {
		logger.Error("Failed in getting the tiny url ", tinyURLRegex)
		return " ", errors.New("Unable to generate tiny URL")
	}
	tinyURL := tinyURLData[startIndex : startIndex+6]
	if _, ok := UrlMapping[longURL]; !ok {
		UrlMapping[longURL] = tinyURL
		repository.WriteFile(FilePath, longURL, tinyURL)

	}

	return tinyURL, nil
}

func LoadMap(filePath string) error {
	logger.Info("In LoadMap Function,filePath is: ", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		logger.Error("Error in opening the file is: ", err)
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		UrlMapping[data[0]] = data[1]
	}
	if err := scanner.Err(); err != nil {
		logger.Error("Error in scanning the data is: ", err)
		return err
	}
	return nil
}

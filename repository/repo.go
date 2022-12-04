package repository

import (
	"os"

	logger "github.com/sirupsen/logrus"
)

func WriteFile(fileName string, longUrl string, tinyUrl string) error {
	logger.Infof("In WriteFile Function. Url is %s and tinyUrl is %s ", longUrl, tinyUrl)
	f, errOpen := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
	if errOpen != nil {
		logger.Error("Error in opening a file is: ", errOpen)
		return errOpen
	}
	defer f.Close()
	_, errWrite := f.WriteString(longUrl + "," + tinyUrl + "\n")
	if errWrite != nil {
		logger.Error("Error in writting  a file is: ", errWrite)
		return errWrite
	}
	return nil
}

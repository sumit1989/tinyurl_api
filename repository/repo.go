package repository

import (
	"os"

	logger "github.com/sirupsen/logrus"
)

type fileCreate struct {
	fileCreat *os.File
}

func NewFile(f *os.File) fileCreate {
	return fileCreate{
		fileCreat: f,
	}
}

func WriteFile(longUrl string, tinyUrl string) {
	logger.Infof("In WriteFile Function. Url is %s and tinyUrl is %s ", longUrl, tinyUrl)
	f, errOpen := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if errOpen != nil {
		logger.Error("Error in opening a file is: ", errOpen)
	}
	defer f.Close()
	_, errWrite := f.WriteString(longUrl + "," + tinyUrl + "\n")
	if errWrite != nil {
		logger.Error("Error in writting  a file is: ", errWrite)
	}

}

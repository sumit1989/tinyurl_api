package main

import (
	"fmt"
	"net/http"
	conf "tinyurl_api/configuration"
	router "tinyurl_api/router"
	"tinyurl_api/utility"

	logger "github.com/sirupsen/logrus"
)

func main() {
	conf, err := conf.Configuration("configuration/config.json")
	if err != nil {
		logger.Error("Not able to load the config.json", err)

	}
	errLoadMap := utility.LoadMap(utility.FilePath)
	if errLoadMap != nil {
		logger.Error("Not able to load the tinyurl map", err)

	}
	ginRouter := router.Setup(conf)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%s", conf.Port), ginRouter); err != nil && err != http.ErrServerClosed {
		logger.Error("Not able to connect to the server", err)
	}

}

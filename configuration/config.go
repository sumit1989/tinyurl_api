package config

import (
	"encoding/json"
	"os"

	logger "github.com/sirupsen/logrus"
)

type Config struct {
	Port string `json:"port"`
}

func Configuration(filepath string) (Config, error) {
	logger.Info("In Configuration Function, filepath is: ", filepath)
	file, err := os.Open(filepath)
	if err != nil {
		logger.Error("Error in opening the file is: ", err)
		return Config{}, err
	}
	var config Config
	d := json.NewDecoder(file)
	decodeErr := d.Decode(&config)
	if decodeErr != nil {
		logger.Error("Error in decoding  the data is: ", decodeErr)
		return Config{}, decodeErr
	}
	return config, nil

}

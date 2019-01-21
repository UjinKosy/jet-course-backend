package main

import (
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
)

// AppConfig contains app's configuration
type AppConfig struct {
	DB     string
	Server struct {
		Path string
		Port string
	}

	Debug bool
}

// Default contains global app's configuration
var DefaultConfig AppConfig

//Load method loads and parses config file
func (c *AppConfig) Load(url string) {
	err := configor.Load(c, url)
	if err != nil {
		log.Error("Log file not found: " + url)
	}
}

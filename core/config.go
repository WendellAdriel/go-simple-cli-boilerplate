package core

import (
	"log"

	"gopkg.in/ini.v1"
)

const iniPath = "app.ini"

type config struct {
	cfg *ini.File
}

// InitConfig will load the application configuration
func (c *config) InitConfig() {
	config, err := ini.InsensitiveLoad(iniPath)
	if err != nil {
		log.Fatal("Error while reading the config file: " + err.Error())
	}
	c.cfg = config
}

// GetString gets a string configuration value
func (c *config) GetString(section, key string) string {
	return c.cfg.Section(section).Key(key).String()
}

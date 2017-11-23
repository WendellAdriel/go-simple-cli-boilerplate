package core

// AppConfig is a wrapper for the application configuration
var AppConfig config

func init() {
	AppConfig.InitConfig()
}

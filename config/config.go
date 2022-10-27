package config

import (
	"os"
	"sync"
)

type AppConfig struct {
	App struct {
		Port string `toml:"port"`
	} `toml:"app"`
	Database struct {
		DBURL  string
		DBNAME string
		DBPASS string
		DBUSER string
		DBETC  string
		DBPORT string
	} `toml:"database"`
	Secrettoken struct {
		Token string `toml:"token"`
	} `toml:"secrettoken"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}
	return appConfig
}

func initConfig() *AppConfig {
	var finalConfig AppConfig
	finalConfig.App.Port = "8080"
	finalConfig.Database.DBNAME = os.Getenv("MONGO_DBNAME")
	finalConfig.Database.DBURL = os.Getenv("MONGO_HOST")
	finalConfig.Database.DBPASS = os.Getenv("MONGO_PASS")
	finalConfig.Database.DBUSER = os.Getenv("MONGO_USER")
	finalConfig.Database.DBETC = os.Getenv("MONGO_ETC")
	finalConfig.Database.DBPORT = os.Getenv("MONGO_PORT")

	return &finalConfig
}

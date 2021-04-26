package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	AppPath string
)

func Start(path string) {
	//set default timezone
	os.Setenv("TZ", "Asia/Jakarta")

	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err.Error()))
	}

	// Open DB Connection
	OpenDB()

	// Open Redis Connection
	OpenRedis()
}

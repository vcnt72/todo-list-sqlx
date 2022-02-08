package config

import (
	"github.com/spf13/viper"
)

// Init is config initializer to setup config file
func Init() {

	// Finding config.yaml on config folder
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

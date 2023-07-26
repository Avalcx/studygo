package config

import "github.com/spf13/viper"

var Address string

func setHttpServerConfig() {
	viper.SetDefault("PORT", "8888")
	viper.AutomaticEnv()
	port := viper.GetString("PORT")
	Address = "0.0.0.0:" + port
}

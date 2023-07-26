package config

import (
	"project01/utils/md5util"

	"github.com/spf13/viper"
)

var (
	Username string
	Password string
)

func setAuthConfig() {
	username := "ava"
	password := "123456"
	viper.SetDefault("USER", username)
	viper.SetDefault("PASSWD", md5util.CreateSecret(password))
	viper.AutomaticEnv()
	Username = viper.GetString("USER")
	Password = viper.GetString("PASSWD")
}

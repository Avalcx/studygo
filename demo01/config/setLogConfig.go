package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func setLogConfig() {
	viper.SetDefault("LOGLEVEL", "debug")
	viper.SetDefault("LOGREPORTER", "false")
	viper.AutomaticEnv()
	level := viper.GetString("LOGLEVEL")
	reporter := viper.GetString("LOGREPORTER")

	if level == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else if level == "error" {
		logrus.SetLevel(logrus.ErrorLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	if reporter == "true" {
		logrus.SetReportCaller(true)
	} else {
		logrus.SetReportCaller(false)
	}
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
}

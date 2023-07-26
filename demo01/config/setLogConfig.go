package config

import (
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func setLogConfig() {
	viper.SetDefault("LOGLEVEL", "debug")
	viper.SetDefault("LOGREPORTER", "true")
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
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",

		// 获取函数所在文件的文件名
		// f.Line 行号 f.File 文件名 f.Function 函数名
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			fileName := path.Base(f.File)
			line := f.Line
			fileNameLine := fileName + ":" + strconv.Itoa(line)
			return f.Function, fileNameLine
		},
	})
}

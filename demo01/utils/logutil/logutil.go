package logutil

import (
	"github.com/sirupsen/logrus"
)

func Debug(fields map[string]any, msg any) {
	logrus.WithFields(fields).Debug(msg)
}

func Info(fields map[string]any, msg any) {
	logrus.WithFields(fields).Info(msg)
}

func Error(fields map[string]any, msg any) {
	logrus.WithFields(fields).Error(msg)
}

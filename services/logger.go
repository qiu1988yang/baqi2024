package services

import "github.com/sirupsen/logrus"

var log *logrus.Logger

func init() {
	log = logrus.New()
	// 这里可以根据需求进行日志格式、级别等的配置
}

func Log() *logrus.Logger {
	return log
}

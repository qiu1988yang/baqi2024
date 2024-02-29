package services

import (
	"log"
	"os"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Info(filename, msg string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("无法打开日志文件：%v", err)
	}
	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Printf("[INFO] %s", msg)
}

func (l *Logger) Warning(filename, msg string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("无法打开日志文件：%v", err)
	}
	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Printf("[WARNING] %s", msg)
}

func (l *Logger) Error(filename, msg string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("无法打开日志文件：%v", err)
	}
	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Printf("[ERROR] %s", msg)
}

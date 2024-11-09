package config

import (
    "github.com/sirupsen/logrus"
    "os"
)

var Port string

func Setup() {
    // Настроим конфигурации
    Port = ":8080" // Например, указываем порт для сервера

    // Настройка логгера
    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.SetOutput(os.Stdout)
    logrus.SetLevel(logrus.InfoLevel)
}

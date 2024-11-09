package config

import (
    "github.com/sirupsen/logrus"
    "os"
)

var Host string
var Port string

func Setup() {
    // Получаем хост из переменной окружения или используем значение по умолчанию
    Host = os.Getenv("HOST")
    if Host == "" {
        Host = "0.0.0.0"
    }

    // Указываем порт для сервера, также можно задать через переменную окружения
    Port = os.Getenv("PORT")
    if Port == "" {
        Port = "8080"
    }

    // Настройка логгера
    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.SetOutput(os.Stdout)
    logrus.SetLevel(logrus.InfoLevel)
}

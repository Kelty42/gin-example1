package main

import (
    "github.com/gin-gonic/gin"
    "github.com/Kelty42/gin-example1/handlers"
    "github.com/Kelty42/gin-example1/middlewares"
    "github.com/Kelty42/gin-example1/config"
)

func main() {
    // Настройки конфигурации (например, логгера, порта и др.)
    config.Setup()

    // Создаем экземпляр Gin
    r := gin.New()

    // Добавляем middleware для логирования
    r.Use(middlewares.JSONLoggerMiddleware())

    // Маршруты
    r.GET("/metrics", handlers.MetricsHandler)
    r.GET("/api", handlers.APIHandler)
    r.GET("/test", handlers.TestHandler)

    // Обработчик для ошибок 404
    r.NoRoute(handlers.NoRouteHandler)

    // Запускаем сервер
    r.Run(config.Port)
}

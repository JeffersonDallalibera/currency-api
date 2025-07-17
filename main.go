package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"       // Swagger middleware
	"github.com/swaggo/files"             // Swagger UI files
	_ "github.com/JeffersonDallalibera/currency-api/docs" // Docs gerados pelo swag

	"github.com/JeffersonDallalibera/currency-api/handler"
	"github.com/JeffersonDallalibera/currency-api/config"
)

// @title API de Cotação de Moedas
// @version 1.0
// @description API para converter valores entre moedas usando exchangerate.host
// @contact.name Jefferson Dallalibera
// @contact.url https://github.com/JeffersonDallalibera
// @host localhost:8080
// @BasePath /
func main() {
	config.Load() 
	r := gin.Default()

	r.GET("/convert", handler.ConvertCurrency)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

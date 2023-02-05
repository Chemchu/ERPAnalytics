package main

import (
	"log"

	endpoints "github.com/Chemchu/ERPAnalytics/endpoints_handlers"
	endpoints_productos "github.com/Chemchu/ERPAnalytics/endpoints_handlers/productos"
	endpoints_ventas "github.com/Chemchu/ERPAnalytics/endpoints_handlers/ventas"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Some error occured. Err: %s", errEnv)
	}

	router := gin.Default()
	router.GET("/", endpoints.GetAPI)
	router.GET("/api", endpoints.GetAPI)
	router.POST("/api/analytics/ventas/summary", endpoints_ventas.PostAnalyzeSales)
	router.POST("/api/analytics/productos/summary", endpoints_productos.PostAnalyzeProducts)

	router.Run("0.0.0.0:6060")
}

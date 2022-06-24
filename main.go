package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Bienvenido al API de Análisis de datos de ERPSolution", "successful": true})
}

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Some error occured. Err: %s", errEnv)
	}

	router := gin.Default()
	router.GET("/api", getAPI)
	// router.GET("/api/analitycs/*object", getStats)

	router.Run("0.0.0.0:6060")
}

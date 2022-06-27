package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Chemchu/ERPAnalytics/analitycs"
	"github.com/Chemchu/ERPAnalytics/types"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Bienvenido al API de Análisis de datos de ERPSolution", "successful": true})
}

func postSummarizeSales(c *gin.Context) {
	var ventasPointer []types.Venta
	if err := c.ShouldBindJSON(&ventasPointer); err != nil {
		fmt.Printf("Error: %+v", err)
		c.JSON(http.StatusOK, gin.H{"message": err, "successful": false})
	}
	summaryResponse := analitycs.GetSalesSummaryByDay(&ventasPointer, "a")
	c.JSON(http.StatusOK, summaryResponse)
}

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Some error occured. Err: %s", errEnv)
	}

	router := gin.Default()
	router.GET("/api", getAPI)
	router.POST("/api/analytics/summary", postSummarizeSales)

	router.Run("0.0.0.0:6060")
}

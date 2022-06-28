package main

import (
	"fmt"
	"log"
	"net/http"

	analitycs "github.com/Chemchu/ERPAnalytics/analytics"
	"github.com/Chemchu/ERPAnalytics/types"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Bienvenido al API de Análisis de datos de ERPSolution", "successful": true})
}

func postSummarizeSales(c *gin.Context) {
	var ventas []types.Venta
	if err := c.ShouldBindJSON(&ventas); err != nil {
		fmt.Printf("Error: %+v\n", err)
		c.JSON(http.StatusOK, gin.H{"message": err, "successful": false})
		return
	}

	summaryResponse := analitycs.GetSalesSummaryByDay(ventas)
	if summaryResponse.Successful {
		c.JSON(http.StatusOK, summaryResponse)
	} else {
		c.JSON(http.StatusBadRequest, summaryResponse)
	}
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

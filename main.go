package main

import (
	"fmt"
	"log"
	"net/http"

	products_analytics "github.com/Chemchu/ERPAnalytics/products_analytics"
	sales_analitycs "github.com/Chemchu/ERPAnalytics/sales_analytics"
	"github.com/Chemchu/ERPAnalytics/types"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Bienvenido al API de Análisis de datos de ERPSolution", "successful": true})
}

func postAnalyzeSales(c *gin.Context) {
	var ventas []types.Venta
	if err := c.ShouldBindJSON(&ventas); err != nil {
		fmt.Printf("Error: %+v\n", err)
		c.JSON(http.StatusOK, gin.H{"message": err, "successful": false})
		return
	}

	summaryResponse := sales_analitycs.GetSalesSummaryByDay(ventas)
	if summaryResponse.Successful {
		c.JSON(http.StatusOK, summaryResponse)
	} else {
		c.JSON(http.StatusBadRequest, summaryResponse)
	}
}

func postAnalyzeProducts(c *gin.Context) {
	var ventas []types.Venta
	if err := c.ShouldBindJSON(&ventas); err != nil {
		fmt.Printf("Error: %+v\n", err)
		c.JSON(http.StatusOK, gin.H{"message": err, "successful": false})
		return
	}

	summaryResponse := products_analytics.GetProductsSummary(ventas)
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
	router.POST("/api/analytics/summary", postAnalyzeSales)

	router.Run("0.0.0.0:6060")
}

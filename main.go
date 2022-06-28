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
	// body, _ := ioutil.ReadAll(c.Request.Body)
	// println(string(body))
	var ventasPointer []types.Venta
	if err := c.ShouldBindJSON(&ventasPointer); err != nil {
		fmt.Printf("Error: %+v\n", err)
		c.JSON(http.StatusOK, gin.H{"message": err, "successful": false})
		return
	}
	// fmt.Println(*ventasPointer[0].ID)
	summaryResponse := analitycs.GetSalesSummaryByDay(&ventasPointer)
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

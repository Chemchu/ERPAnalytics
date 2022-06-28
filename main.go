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

// type Prueba struct {
// 	Msg string `json:"msg"`
// 	//Data Prueba2 `json:"data,omitempty"`
// }

// type Prueba2 struct {
// 	Succ bool `json:"msg"`
// }

func postSummarizeSales(c *gin.Context) {
	var ventas []types.Venta
	if err := c.ShouldBindJSON(&ventas); err != nil {
		fmt.Printf("Error: %+v\n", err)
		c.JSON(http.StatusOK, gin.H{"message": err, "successful": false})
		return
	}

	// p2 := Prueba2{
	// 	Succ: true,
	// }
	// p1 := Prueba{
	// 	Msg: "prueba1",
	// 	//Data: p2,
	// }
	// pRes, _ := json.Marshal(&p1)

	summaryResponse := analitycs.GetSalesSummaryByDay(ventas)
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

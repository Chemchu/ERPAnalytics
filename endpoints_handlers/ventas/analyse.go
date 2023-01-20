package endpoints_ventas

import (
	"fmt"
	"net/http"

	sales_analitycs "github.com/Chemchu/ERPAnalytics/sales_analytics"
	"github.com/Chemchu/ERPAnalytics/types"
	"github.com/gin-gonic/gin"
)

func PostAnalyzeSales(c *gin.Context) {
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

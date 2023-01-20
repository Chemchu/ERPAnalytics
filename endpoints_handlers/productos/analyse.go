package endpoints_productos

import (
	"fmt"
	"net/http"

	"github.com/Chemchu/ERPAnalytics/products_analytics"
	"github.com/Chemchu/ERPAnalytics/types"
	"github.com/gin-gonic/gin"
)

func PostAnalyzeProducts(c *gin.Context) {
	var ventas []types.Venta
	var productosIds []string
	if err := c.ShouldBindJSON(&ventas); err != nil {
		fmt.Printf("Error: %+v\n", err)
		c.JSON(http.StatusOK, gin.H{"message": err, "successful": false})
		return
	}

	if err := c.ShouldBindJSON(&productosIds); err != nil {
		fmt.Printf("Error: %+v\n", err)
		c.JSON(http.StatusOK, gin.H{"message": err, "successful": false})
		return
	}

	summaryResponse := products_analytics.GetProductsSummary(ventas, productosIds)
	if summaryResponse.Successful {
		c.JSON(http.StatusOK, summaryResponse)
	} else {
		c.JSON(http.StatusBadRequest, summaryResponse)
	}
}

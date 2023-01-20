package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Bienvenido al API de Análisis de datos de ERPSolution", "successful": true})
}

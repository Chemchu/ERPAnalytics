package products_analytics

import (
	"encoding/json"
	"fmt"
	// "math"
	// "sort"
	// "strconv"
	// "time"
	//
	// "github.com/Chemchu/ERPAnalytics/placeholders"
	"github.com/Chemchu/ERPAnalytics/types"
)

func GetProductsSummary(ventas []types.Venta) types.APIResponse {
	msg := "Petición realizada correctamente"
	successful := true
	data, err := json.Marshal(Summarize(&ventas))
	if err != nil {
		msg = fmt.Sprintf("Error al convertir el Summary a JSON: %s", err.Error())
		successful = false
		data = nil
	}

	return types.APIResponse{
		Message:    msg,
		Successful: successful,
		Data:       string(data),
	}
}

func Summarize(ventas *[]types.Venta) types.ProductsSummary {
	if len(*ventas) <= 0 {
		return types.ProductsSummary{}
	}
	return types.ProductsSummary{}
}

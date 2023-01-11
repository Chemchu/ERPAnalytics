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

func Summarize(ventas *[]types.Venta) []types.ProductSummary {
	if len(*ventas) <= 0 {
		return []types.ProductSummary{}
	}

	// Map de los productosMap vendidos
	productosMap := make(map[string]types.ProductSummary)

	// Iterar sobre cada producto vendido
	for i := 0; i < len(*ventas); i++ {
		productos := (*ventas)[i].Productos
		for j := 0; j < len(productos); j++ {
			productoSummary, existeSummary := productosMap[productos[j].Ean]
			producto := productos[j]

			if !existeSummary {
				var summary = types.ProductSummary{
					IDProducto:         producto.ID,
					NombreProducto:     producto.Nombre,
					Ean:                producto.Ean,
					Familia:            producto.Familia,
					Proveedor:          producto.Proveedor,
					CantidadVendida:    int32(producto.CantidadVendida),
					CosteTotalProducto: float64(producto.CantidadVendida) * float64(producto.PrecioCompra),
				}
				// Actualizar el MAP
				productosMap[producto.Ean] = summary
			} else {
				// Actualizar el summary con el nuevo producto vendido
				updatedSummary := types.ProductSummary{}

				productosMap[producto.Ean] = updatedSummary
			}
		}
	}

	return []types.ProductSummary{}
}

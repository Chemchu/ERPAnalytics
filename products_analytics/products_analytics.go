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

			precioConIva := producto.PrecioCompra + (producto.PrecioCompra * (producto.Iva / 100))
			beneficioProducto := producto.PrecioFinal - precioConIva
			ivaProducto := producto.PrecioCompra * (producto.Iva / 100)

			// Si el producto todavía no se ha añadido al Map
			if !existeSummary {
				var summary = types.ProductSummary{
					IDProducto:         producto.ID,
					NombreProducto:     producto.Nombre,
					Ean:                producto.Ean,
					Familia:            producto.Familia,
					Proveedor:          producto.Proveedor,
					CantidadVendida:    int32(producto.CantidadVendida),
					CosteTotalProducto: float64(producto.CantidadVendida) * float64(producto.PrecioCompra),
					VentaTotal:         producto.PrecioFinal * float64(producto.CantidadVendida),
					Beneficio:          float64(producto.CantidadVendida) * beneficioProducto,
					IVAPagado:          float64(producto.CantidadVendida) * ivaProducto,
				}
				// Actualizar el MAP
				productosMap[producto.Ean] = summary
			} else {
				// En caso de que el producto exista en el Map, actualizar el summary
				var updatedSummary = types.ProductSummary{
					IDProducto:         producto.ID,
					NombreProducto:     producto.Nombre,
					Ean:                producto.Ean,
					Familia:            producto.Familia,
					Proveedor:          producto.Proveedor,
					CantidadVendida:    productoSummary.CantidadVendida + int32(producto.CantidadVendida),
					CosteTotalProducto: productoSummary.CosteTotalProducto + (float64(producto.CantidadVendida) * float64(producto.PrecioCompra)),
					VentaTotal:         productoSummary.VentaTotal + (producto.PrecioFinal * float64(producto.CantidadVendida)),
					Beneficio:          productoSummary.Beneficio + (float64(producto.CantidadVendida) * beneficioProducto),
					IVAPagado:          productoSummary.IVAPagado + (float64(producto.CantidadVendida) * ivaProducto),
				}
				productosMap[producto.Ean] = updatedSummary
			}
		}
	}

	return []types.ProductSummary{}
}

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
	"github.com/Chemchu/ERPAnalytics/data_structure"
	"github.com/Chemchu/ERPAnalytics/types"
	"github.com/golang-module/carbon/v2"
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

	productosMap := data_structure.StringProductSummaryMap{}
	fechasSet := data_structure.StringSet{}

	for i := 0; i < len(*ventas); i++ {
		fecha := carbon.CreateFromTimestampMilli((*ventas)[i].CreatedAt).ToDateString()
		if !fechasSet.Has(fecha) {
			fechasSet.Add(fecha)
		}

		productos := (*ventas)[i].Productos
		for j := 0; j < len(productos); j++ {
			productoSummary, existeSummary := productosMap.Has(productos[j].Ean)
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

				productosMap.Add(producto.Ean, summary)
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

				productosMap.Add(producto.Ean, updatedSummary)
			}
		}
	}

	for ean, producto := range productosMap {
		producto.FrecuenciaVentaDiaria = float64(producto.CantidadVendida) / float64(fechasSet.Length())
		productosMap.Add(ean, producto)
	}

	return productosMap.Values()
}

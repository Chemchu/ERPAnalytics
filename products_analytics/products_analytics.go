package products_analytics

import (
	"encoding/json"
	"fmt"

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
		productosDeLaVenta := (*ventas)[i].Productos
		fecha := carbon.CreateFromTimestampMilli((*ventas)[i].CreatedAt).ToDateString()
		if !fechasSet.Has(fecha) {
			fechasSet.Add(fecha)
		}

		for j := 0; j < len(productosDeLaVenta); j++ {
			productoSummary, existeSummary := productosMap.Has(productosDeLaVenta[j].Ean)
			productoVendido := productosDeLaVenta[j]

			precioConIva := productoVendido.PrecioCompra + (productoVendido.PrecioCompra * (productoVendido.Iva / 100))
			beneficioProducto := productoVendido.PrecioFinal - precioConIva
			ivaProducto := productoVendido.PrecioCompra * (productoVendido.Iva / 100)

			// Si el producto todavía no se ha añadido al Map
			if !existeSummary {
				summary := types.ProductSummary{
					IDProducto:         productoVendido.ID,
					NombreProducto:     productoVendido.Nombre,
					Ean:                productoVendido.Ean,
					Familia:            productoVendido.Familia,
					Proveedor:          productoVendido.Proveedor,
					CantidadVendida:    int32(productoVendido.CantidadVendida),
					CosteTotalProducto: float64(productoVendido.CantidadVendida) * float64(productoVendido.PrecioCompra),
					VentaTotal:         productoVendido.PrecioFinal * float64(productoVendido.CantidadVendida),
					Beneficio:          float64(productoVendido.CantidadVendida) * beneficioProducto,
					IVAPagado:          float64(productoVendido.CantidadVendida) * ivaProducto,
				}

				productosMap.Add(productoVendido.Ean, summary)
			} else {
				// En caso de que el producto exista en el Map, actualizar el summary
				productoSummary.AddCantidadvendida(int32(productoVendido.CantidadVendida))
				productoSummary.AddCosteTotalProducto(float64(productoVendido.CantidadVendida) * float64(productoVendido.PrecioCompra))
				productoSummary.AddValorVenta(float64(productoVendido.CantidadVendida) * productoVendido.PrecioFinal)
				productoSummary.AddBeneficio(float64(productoVendido.CantidadVendida) * beneficioProducto)
				productoSummary.AddIVAPagado(float64(productoVendido.CantidadVendida) * ivaProducto)

				productosMap.Add(productoVendido.Ean, productoSummary)
			}
		}
	}

	productos := AddFrecuenciaVentaDiaria(&productosMap, &fechasSet)
	return productos
}

func AddFrecuenciaVentaDiaria(productosMap *data_structure.StringProductSummaryMap, fechasSet *data_structure.StringSet) []types.ProductSummary {
	for ean, producto := range *productosMap {
		producto.FrecuenciaVentaDiaria = float64(producto.CantidadVendida) / float64(fechasSet.Length())
		productosMap.Add(ean, producto)
	}
	return productosMap.Values()
}

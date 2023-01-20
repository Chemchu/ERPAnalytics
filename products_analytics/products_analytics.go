package products_analytics

import (
	"encoding/json"
	"fmt"

	"github.com/Chemchu/ERPAnalytics/data_structure"
	"github.com/Chemchu/ERPAnalytics/types"
	"github.com/golang-module/carbon/v2"
)

func GetProductsSummary(ventas []types.Venta, productosIds []string) types.APIResponse {
	if len(ventas) <= 0 {
		return types.APIResponse{
			Message:    "No existen ventas para las fechas seleccionadas",
			Successful: true,
			Data:       "",
		}
	}
	if len(productosIds) <= 0 {
		return types.APIResponse{
			Message:    "Es necesario seleccionar al menos un producto para realizar el analisis",
			Successful: true,
			Data:       "",
		}
	}

	productos := data_structure.StringSet{}
	for i := 0; i < len(productosIds); i++ {
		if !productos.Has(productosIds[i]) {
			productos.Add(productosIds[i])
		}
	}

	msg := "Petición realizada correctamente"
	successful := true
	data, err := json.Marshal(Summarize(&ventas, &productos))
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

func Summarize(ventas *[]types.Venta, productosIds *data_structure.StringSet) []types.ProductoSummary {
	if len(*ventas) <= 0 {
		return []types.ProductoSummary{}
	}
	if len(*productosIds) <= 0 {
		return []types.ProductoSummary{}
	}

	productos := data_structure.StringProductSummaryMap{}
	fechas := data_structure.StringSet{}

	for i := 0; i < len(*ventas); i++ {
		productosVenta := (*ventas)[i].Productos
		fecha := carbon.CreateFromTimestampMilli((*ventas)[i].CreatedAt).ToDateString()
		if !fechas.Has(fecha) {
			fechas.Add(fecha)
		}

		for j := 0; j < len(productosVenta); j++ {
			productoVendido := productosVenta[j]
			if !productosIds.Has(productoVendido.ID) {
				continue
			}

			productoSummary, existeSummary := productos.Has(productoVendido.ID)
			if !existeSummary {
				summary := types.ProductoSummary{
					IDProducto:     productoVendido.ID,
					NombreProducto: productoVendido.Nombre,
					Ean:            productoVendido.Ean,
					Familia:        productoVendido.Familia,
					Proveedor:      productoVendido.Proveedor,
				}

				productos.Add(productoVendido.ID, summary.UpdateSummary(productoVendido))
				continue
			}

			productos.Add(productoVendido.ID, productoSummary.UpdateSummary(productoVendido))
		}
	}

	productos.UpdateFrecuenciaVentaDiaria(&fechas)
	return productos.Values()
}

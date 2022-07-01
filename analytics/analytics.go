package analitycs

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/Chemchu/ERPAnalytics/types"
)

func GetSalesSummaryByDay(ventas []types.Venta) types.APIResponse {
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

func MapToArray(mapObject map[string]types.VentasPorHora) []types.VentasPorHora {
	horasSorted := make([]string, 0, len(mapObject))
	for k := range mapObject {
		horasSorted = append(horasSorted, k)
	}
	sort.Strings(horasSorted)

	resArray := make([]types.VentasPorHora, 0, len(mapObject))
	for _, val := range horasSorted {
		resArray = append(resArray, mapObject[val])
	}

	return resArray
}

func GetMostFrequentValue(array []string) string {
	// Collect words and their count in a histogram.
	histo := make(map[string]int)
	for _, str := range array {
		histo[str]++
	}

	// Scan the histogram to find the word with the highest count.
	frecuencia := 0
	palabraMasFrecuente := ""
	for key, value := range histo {
		if value > frecuencia {
			frecuencia = value
			palabraMasFrecuente = key
		}
	}
	return palabraMasFrecuente
}

func Summarize(ventas *[]types.Venta) types.Summary {
	if len(*ventas) <= 0 {
		return types.Summary{
			VentasPorHora:             []types.VentasPorHora{},
			Beneficio:                 0.0,
			TotalVentas:               0.0,
			NumVentas:                 0,
			DineroDescontado:          0.0,
			CantidadProductosVendidos: 0,
			MediaVentas:               0,
			MediaCantidadVenida:       0,
			IVAPagado:                 0.0,
			TotalEfectivo:             0.0,
			TotalTarjeta:              0.0,
		}
	}

	beneficioTotal := 0.0
	dineroDescontadoTotal := 0.0
	total := 0.0
	totalTarjeta := 0.0
	totalEfectivo := 0.0
	ivaPagado := 0.0
	prodVendidosTotal := 0
	ventasPorHoraMap := make(map[string]types.VentasPorHora)
	numVentas := 0
	mediaVentas := 0.0
	mediaCantidadVendida := 0.0

	for _, venta := range *ventas {
		beneficioHora := 0.0
		dineroDescontadoHora := 0.0
		totalHora := 0.0
		totalTarjetaHora := 0.0
		totalEfectivoHora := 0.0
		prodVendidosHora := 0
		hora := strconv.FormatInt(int64(time.UnixMilli(venta.CreatedAt).Hour()), 10)
		FormatHour(&hora)

		// Comprueba si ya hay ventas añadidas para una deterrminada hora
		if ventaEnMap, containsValue := ventasPorHoraMap[hora]; containsValue {
			beneficioHora = ventaEnMap.BeneficioHora
			totalHora = ventaEnMap.TotalVentaHora
			totalTarjetaHora = ventaEnMap.TotalTarjetaHora
			totalEfectivoHora = ventaEnMap.TotalEfectivoHora
			prodVendidosHora = ventaEnMap.ProductosVendidosHora
			dineroDescontadoHora = ventaEnMap.DineroDescontadoHora
		}

		if venta.Tipo == "Efectivo" || venta.Tipo == "Cobro rápido" {
			totalEfectivo += venta.PrecioVentaTotal
			totalEfectivoHora += venta.PrecioVentaTotal
		} else {
			if venta.Tipo == "Tarjeta" {
				totalTarjeta += venta.PrecioVentaTotal
				totalTarjetaHora += venta.PrecioVentaTotal
			} else {
				totalEfectivo += venta.PrecioVentaTotal - venta.Cambio
				totalTarjeta += venta.PrecioVentaTotal - venta.Cambio

				totalEfectivoHora += venta.DineroEntregadoEfectivo - venta.Cambio
				totalTarjetaHora += venta.DineroEntregadoTarjeta
			}
		}

		for _, producto := range venta.Productos {
			beneficioTotal += (producto.PrecioFinal - producto.PrecioCompra) * float64(producto.CantidadVendida)
			beneficioHora += (producto.PrecioFinal - producto.PrecioCompra) * float64(producto.CantidadVendida)
			prodVendidosTotal += producto.CantidadVendida
			prodVendidosHora += producto.CantidadVendida
			dineroDescontadoHora += (producto.PrecioVenta - producto.PrecioFinal) * float64(producto.CantidadVendida)
			ivaPagado += producto.PrecioCompra * (producto.Iva / 100)
		}

		total += venta.PrecioVentaTotal
		totalHora += venta.PrecioVentaTotal

		ventaPorHora := types.VentasPorHora{
			Hora:                  hora,
			BeneficioHora:         beneficioHora,
			TotalVentaHora:        totalHora,
			TotalEfectivoHora:     totalEfectivoHora,
			TotalTarjetaHora:      totalTarjetaHora,
			ProductosVendidosHora: prodVendidosHora,
			DineroDescontadoHora:  dineroDescontadoHora,
		}
		ventasPorHoraMap[hora] = ventaPorHora
	}

	ventasPorHora := MapToArray(ventasPorHoraMap)
	numVentas = len(*ventas)
	mediaVentas = total / float64(numVentas)
	if numVentas > 0 {
		mediaCantidadVendida = float64(prodVendidosTotal / numVentas)
	}

	return types.Summary{
		VentasPorHora:             ventasPorHora,
		Beneficio:                 beneficioTotal,
		TotalVentas:               total,
		TotalEfectivo:             totalEfectivo,
		TotalTarjeta:              totalTarjeta,
		NumVentas:                 numVentas,
		DineroDescontado:          dineroDescontadoTotal,
		CantidadProductosVendidos: prodVendidosTotal,
		MediaVentas:               mediaVentas,
		MediaCantidadVenida:       mediaCantidadVendida,
		IVAPagado:                 ivaPagado,
	}
}

func FormatHour(hora *string) {
	if len(*hora) < 2 {
		*hora = "0" + *hora
	}
	*hora += ":00"
}

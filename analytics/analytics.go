package analitycs

import (
	"encoding/json"
	"fmt"
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

func MapToArray(d map[string]types.VentasPorHora) []types.VentasPorHora {
	m := make([]types.VentasPorHora, 0, len(d))
	for _, val := range d {
		m = append(m, val)
	}

	return m
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
	beneficioTotal := 0.0
	dineroDescontadoTotal := 0.0
	total := 0.0
	totalTarjeta := 0.0
	totalEfectivo := 0.0
	ivaPagado := 0.0
	prodVendidosTotal := 0
	ventasPorHoraMap := make(map[string]types.VentasPorHora)
	ventasPorHora := MapToArray(ventasPorHoraMap)
	numVentas := 0
	mediaVentas := 0.0
	mediaCantidadVendida := 0.0

	if len(*ventas) <= 0 {
		return types.Summary{
			VentasPorHora:             ventasPorHora,
			Beneficio:                 beneficioTotal,
			TotalVentas:               total,
			NumVentas:                 numVentas,
			DineroDescontado:          dineroDescontadoTotal,
			CantidadProductosVendidos: prodVendidosTotal,
			MediaVentas:               mediaVentas,
			MediaCantidadVenida:       mediaCantidadVendida,
			IVAPagado:                 ivaPagado,
		}
	}

	for _, venta := range *ventas {
		// Inicializa los valores
		beneficioHora := 0.0
		dineroDescontadoHora := 0.0
		totalHora := 0.0
		totalTarjetaHora := 0.0
		totalEfectivoHora := 0.0
		prodVendidosHora := 0
		hora := strconv.FormatInt(int64(time.UnixMilli(venta.CreatedAt).Hour()), 10)

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
			total += producto.PrecioFinal * float64(producto.CantidadVendida)
			totalHora += producto.PrecioFinal * float64(producto.CantidadVendida)
			prodVendidosTotal += producto.CantidadVendida
			prodVendidosHora += producto.CantidadVendida
			dineroDescontadoHora += (producto.PrecioVenta - producto.PrecioFinal) * float64(producto.CantidadVendida)
			ivaPagado += producto.PrecioCompra * (producto.Iva / 100)
		}

		venta := types.VentasPorHora{
			Hora:                  hora,
			BeneficioHora:         beneficioHora,
			TotalVentaHora:        totalHora,
			TotalEfectivoHora:     totalEfectivoHora,
			TotalTarjetaHora:      totalTarjetaHora,
			ProductosVendidosHora: prodVendidosHora,
			DineroDescontadoHora:  dineroDescontadoHora,
		}
		ventasPorHoraMap[hora] = venta
	}

	ventasPorHora = MapToArray(ventasPorHoraMap)
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

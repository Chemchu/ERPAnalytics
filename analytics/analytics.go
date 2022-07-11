package analitycs

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"

	"github.com/Chemchu/ERPAnalytics/placeholders"
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
			VentasPorHora:             placeholders.VentasPorHoraPlaceholder(),
			ProductosMasVendidos:      placeholders.ProductosMasVendidosPlaceholder(),
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
	productosVendidosMap := make(map[string]int)
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
		CalcularDineroEntregado(venta, &totalEfectivo, &totalTarjeta, &totalEfectivoHora, &totalTarjetaHora)

		for _, producto := range venta.Productos {
			beneficioTotal += CalcularBeneficio(producto)
			beneficioHora += CalcularBeneficio(producto)
			prodVendidosTotal += producto.CantidadVendida
			prodVendidosHora += producto.CantidadVendida
			dineroDescontadoHora += CalcularDTOAplicado(producto)
			ivaPagado += CalcularIVA(producto)
			productosVendidosMap = UpdateProductosVendidos(producto, productosVendidosMap)
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
		VentasPorHora: ventasPorHora,
		//ProductosMasVendidos:      GetMostFrequent(productosVendidosMap),
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

func CalcularBeneficio(producto types.ProductoVendido) float64 {
	// PrecioFinal = PrecioCompra +* IVA +* Beneficio(Margen)
	// Beneficio = PrecioFinal - PrecioCompra +* IVA
	precioConIva := (producto.PrecioCompra * (producto.Iva / 100)) + producto.PrecioCompra
	beneficio := (producto.PrecioFinal - precioConIva) * float64(producto.CantidadVendida)
	beneficio = math.Round(beneficio*100) / 100

	return beneficio
}

func CalcularIVA(producto types.ProductoVendido) float64 {
	ivaPagado := producto.PrecioCompra * (producto.Iva / 100) * float64(producto.CantidadVendida)
	ivaPagado = math.Round(ivaPagado*100) / 100
	return ivaPagado
}

func CalcularDTOAplicado(producto types.ProductoVendido) float64 {
	dto := (producto.PrecioVenta - producto.PrecioFinal) * float64(producto.CantidadVendida)
	dto = math.Round(dto*100) / 100
	return dto
}

func CalcularDineroEntregado(venta types.Venta, efectivo *float64, tarjeta *float64, efectivoHora *float64, tarjetaHora *float64) {
	if venta.Tipo == "Efectivo" || venta.Tipo == "Cobro rápido" {
		*efectivo += venta.PrecioVentaTotal
		*efectivoHora += venta.PrecioVentaTotal
	} else {
		if venta.Tipo == "Tarjeta" {
			*tarjeta += venta.PrecioVentaTotal
			*tarjetaHora += venta.PrecioVentaTotal
		} else {
			*efectivo += venta.PrecioVentaTotal - venta.Cambio
			*tarjeta += venta.PrecioVentaTotal - venta.Cambio

			*efectivoHora += venta.DineroEntregadoEfectivo - venta.Cambio
			*tarjetaHora += venta.DineroEntregadoTarjeta
		}
	}
}

func UpdateProductosVendidos(producto types.ProductoVendido, productosVendidosMap map[string]int) map[string]int {
	prodCantidad, containValue := productosVendidosMap[producto.Ean]
	if containValue {
		prodCantidad += producto.CantidadVendida
		productosVendidosMap[producto.Ean] = prodCantidad
	} else {
		productosVendidosMap[producto.Ean] = producto.CantidadVendida
	}

	return productosVendidosMap
}

func GetMostFrequent(hMap map[string]int) []string {
	keys := make([]string, 0, len(hMap))
	for key := range hMap {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return hMap[keys[i]] < hMap[keys[j]]
	})

	return keys
}

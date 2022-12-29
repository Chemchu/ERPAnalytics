package sales_analitycs

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

func Summarize(ventas *[]types.Venta) types.SalesSummary {
	if len(*ventas) <= 0 {
		return types.SalesSummary{
			VentasPorHora:             placeholders.VentasPorHoraPlaceholder(),
			ProductosMasVendidos:      []types.ProductoMasVendido{},
			FamiliasMasVendidas:       []types.FamiliaMasVendida{},
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
	ventaMinima := 0.0
	ventaMaxima := 0.0
	mediaCantidadVendida := 0.0
	productosVendidos := []types.ProductoVendido{}

	for _, venta := range *ventas {
		beneficioHora := 0.0
		dineroDescontadoHora := 0.0
		totalHora := 0.0
		totalTarjetaHora := 0.0
		totalEfectivoHora := 0.0
		prodVendidosHora := 0
		hora := strconv.FormatInt(int64(time.UnixMilli(venta.CreatedAt).UTC().Hour()), 10)
		FormatHour(&hora)

		// Comprueba si ya hay ventas añadidas para una determinada hora
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
			productosVendidos = append(productosVendidos, producto)
		}

		total += venta.PrecioVentaTotal
		totalHora += venta.PrecioVentaTotal

		if ventaMinima <= 0.0 || ventaMinima > venta.PrecioVentaTotal {
			ventaMinima = venta.PrecioVentaTotal
		}
		if ventaMaxima <= 0.0 || ventaMaxima < venta.PrecioVentaTotal {
			ventaMaxima = venta.PrecioVentaTotal
		}

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

	prodMasVendidos, familiasMasVendidas := GetMasVendidos(productosVendidos)
	return types.SalesSummary{
		VentasPorHora:             ventasPorHora,
		ProductosMasVendidos:      prodMasVendidos,
		FamiliasMasVendidas:       familiasMasVendidas,
		Beneficio:                 beneficioTotal,
		TotalVentas:               total,
		TotalEfectivo:             totalEfectivo,
		TotalTarjeta:              totalTarjeta,
		NumVentas:                 numVentas,
		DineroDescontado:          dineroDescontadoTotal,
		CantidadProductosVendidos: prodVendidosTotal,
		MediaVentas:               mediaVentas,
		VentaMinima:               ventaMinima,
		VentaMaxima:               ventaMaxima,
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

// Beneficio = precioFinal - precioCompraConIVA
func CalcularBeneficio(producto types.ProductoVendido) float64 {
	precioCompraConIva := producto.PrecioCompra + (producto.PrecioCompra * (producto.Iva / 100))
	beneficio := (producto.PrecioFinal - precioCompraConIva) * float64(producto.CantidadVendida)
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
			*efectivo += venta.DineroEntregadoEfectivo - venta.Cambio
			*tarjeta += venta.DineroEntregadoTarjeta

			*efectivoHora += venta.DineroEntregadoEfectivo - venta.Cambio
			*tarjetaHora += venta.DineroEntregadoTarjeta
		}
	}
}

func GetMasVendidos(productosVendidos []types.ProductoVendido) ([]types.ProductoMasVendido, []types.FamiliaMasVendida) {
	productosVendidosMap := make(map[string]types.ProductoVendido)
	freqVendido := make(map[string]int)
	freqFamilia := make(map[string]int)

	for _, prod := range productosVendidos {
		freqVendido[prod.Ean] += prod.CantidadVendida
		freqFamilia[prod.Familia] += prod.CantidadVendida

		if _, containsProduct := productosVendidosMap[prod.Ean]; !containsProduct {
			productosVendidosMap[prod.Ean] = prod
		}
	}

	familiaMasFrecuentes := []types.FamiliaMasVendida{}
	for familia, cantidad := range freqFamilia {
		familiaMasFrecuentes = append(familiaMasFrecuentes, types.FamiliaMasVendida{
			Familia:         familia,
			CantidadVendida: cantidad,
		})
	}

	productosMasFrecuentes := []types.ProductoMasVendido{}
	for ean, cantidad := range freqVendido {
		productosMasFrecuentes = append(productosMasFrecuentes, types.ProductoMasVendido{
			ID:              productosVendidosMap[ean].ID,
			Nombre:          productosVendidosMap[ean].Nombre,
			Familia:         productosVendidosMap[ean].Familia,
			Ean:             ean,
			CantidadVendida: cantidad,
		})
	}

	sort.Slice(productosMasFrecuentes, func(i, j int) bool {
		return productosMasFrecuentes[i].CantidadVendida > productosMasFrecuentes[j].CantidadVendida
	})

	longitudProd := len(productosMasFrecuentes)
	if longitudProd > 10 {
		longitudProd = 10
	}
	productosMasFrecuentes = productosMasFrecuentes[0:longitudProd]

	longitudFamilia := len(familiaMasFrecuentes)
	if longitudFamilia > 10 {
		longitudFamilia = 10
	}
	familiaMasFrecuentes = familiaMasFrecuentes[0:longitudFamilia]

	return productosMasFrecuentes, familiaMasFrecuentes
}

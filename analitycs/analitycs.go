package analitycs

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Chemchu/ERPAnalytics/types"
)

type APIResponse struct {
	Message    *string  `json:"message"`
	Successful *bool    `json:"successful"`
	Data       *Summary `json:"data"`
}

type Summary struct {
	VentasPorHora   *[]VentaPorHora `json:"ventasPorHora"`
	Fecha           *string         `json:"fecha"`
	Beneficio       *float64        `json:"Beneficio"`
	TotalVentas     *float64        `json:"totalVentas"`
	CantidadVendida *float64        `json:"cantidadVendida"`
}

type VentaPorHora struct {
	Hora              *string  `json:"hora"`
	BeneficioHora     *float64 `json:"beneficioHora"`
	TotalVentaHora    *float64 `json:"totalVentaHora"`
	NumVendido        *int     `json:"cantidadVendida"`
	FamiliaMasVendida *string  `json:"familiaMasVendida"`
}

func GetSalesSummaryByDay(ventas *[]types.Venta, dia string) APIResponse {
	msg := "Petición realizada correctamente"
	successful := true
	ventasPorHoraMap := make(map[string]VentaPorHora)

	for index, venta := range *ventas {
		fmt.Printf("Indice %d: %s", index, *venta.ID)
		hora := strconv.FormatInt(int64(time.UnixMilli(*venta.CreatedAt).Hour()), 10)
		beneficio := 0.0
		total := 0.0
		numVentas := 0
		for _, producto := range *venta.Productos {
			fmt.Println(*producto.PrecioFinal)
			beneficio += *producto.Margen * *producto.PrecioFinal
			total += *producto.PrecioFinal
			numVentas++
		}
		venta := VentaPorHora{
			Hora:           &hora,
			BeneficioHora:  &beneficio,
			TotalVentaHora: &total,
			NumVendido:     &numVentas,
			// FamiliaMasVendida: &fVendida,
		}
		ventasPorHoraMap[hora] = venta
	}

	ventasPorHora := MapToArray(ventasPorHoraMap)
	data := Summary{
		VentasPorHora: &ventasPorHora,
	}
	return APIResponse{
		Message:    &msg,
		Successful: &successful,
		Data:       &data,
	}
}

func MapToArray(d map[string]VentaPorHora) []VentaPorHora {
	m := make([]VentaPorHora, 0, len(d))
	for _, val := range d {
		m = append(m, val)
	}

	return m
}

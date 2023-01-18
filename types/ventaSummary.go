package types

type VentaSummary struct {
	VentasPorHora             []VentasPorHora      `json:"ventasPorHora"`
	ProductosMasVendidos      []ProductoMasVendido `json:"productosMasVendidos"`
	FamiliasMasVendidas       []FamiliaMasVendida  `json:"familiasMasVendidas"`
	Beneficio                 float64              `json:"beneficio"`
	TotalVentas               float64              `json:"totalVentas"`
	TotalEfectivo             float64              `json:"totalEfectivo"`
	TotalTarjeta              float64              `json:"totalTarjeta"`
	NumVentas                 int                  `json:"numVentas"`
	MediaVentas               float64              `json:"mediaVentas"`
	VentaMinima               float64              `json:"ventaMinima"`
	VentaMaxima               float64              `json:"ventaMaxima"`
	MediaCantidadVenida       float64              `json:"mediaCantidadVenida"`
	CantidadProductosVendidos int                  `json:"cantidadProductosVendidos"`
	DineroDescontado          float64              `json:"dineroDescontado"`
	IVAPagado                 float64              `json:"ivaPagado"`
}

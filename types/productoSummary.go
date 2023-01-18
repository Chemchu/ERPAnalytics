package types

type ProductSummary struct {
	IDProducto            string  `json:"_id"`
	NombreProducto        string  `json:"nombreProducto"`
	Ean                   string  `json:"ean"`
	Familia               string  `json:"familia"`
	Proveedor             string  `json:"proveedor"`
	CantidadVendida       int32   `json:"cantidadVendida"`
	CosteTotalProducto    float64 `json:"costeTotalProducto"`
	VentaTotal            float64 `json:"ventaTotal"`
	Beneficio             float64 `json:"beneficio"`
	IVAPagado             float64 `json:"ivaPagado"`
	FrecuenciaVentaDiaria float64 `json:"frecuentaVentaDiaria"` // Cuantas unidades se venden al día
}

func (ps *ProductSummary) AddCantidadvendida(cantidadVendida int32) {
	ps.CantidadVendida += int32(cantidadVendida)
}

func (ps *ProductSummary) AddBeneficio(beneficio float64) {
	ps.Beneficio += float64(beneficio)
}

func (ps *ProductSummary) AddIVAPagado(ivaPagado float64) {
	ps.IVAPagado += float64(ivaPagado)
}

func (ps *ProductSummary) AddCosteTotalProducto(costeProducto float64) {
	ps.CosteTotalProducto += float64(costeProducto)
}

func (ps *ProductSummary) AddValorVenta(valorVenta float64) {
	ps.VentaTotal += float64(valorVenta)
}

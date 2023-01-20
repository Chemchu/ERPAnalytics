package types

type ProductoSummary struct {
	IDProducto            string  `json:"_id"`
	NombreProducto        string  `json:"nombreProducto"`
	Ean                   string  `json:"ean"`
	Familia               string  `json:"familia"`
	Proveedor             string  `json:"proveedor"`
	CantidadVendida       int32   `default0:"0" json:"cantidadVendida"`
	CosteTotalProducto    float64 `default0:"0.0" json:"costeTotalProducto"`
	VentaTotal            float64 `default0:"0.0" json:"ventaTotal"`
	Beneficio             float64 `default0:"0.0" json:"beneficio"`
	IVAPagado             float64 `default0:"0.0" json:"ivaPagado"`
	FrecuenciaVentaDiaria float64 `default0:"0.0" json:"frecuentaVentaDiaria"` // Cuantas unidades se venden al día
}

func (ps *ProductoSummary) AddCantidadvendida(cantidadVendida int32) {
	ps.CantidadVendida += int32(cantidadVendida)
}

func (ps *ProductoSummary) AddBeneficio(beneficio float64) {
	ps.Beneficio += float64(beneficio)
}

func (ps *ProductoSummary) AddIVAPagado(ivaPagado float64) {
	ps.IVAPagado += float64(ivaPagado)
}

func (ps *ProductoSummary) AddCosteTotalProducto(costeProducto float64) {
	ps.CosteTotalProducto += float64(costeProducto)
}

func (ps *ProductoSummary) AddValorVenta(valorVenta float64) {
	ps.VentaTotal += float64(valorVenta)
}

func (ps ProductoSummary) UpdateSummary(productoVendido ProductoVendido) ProductoSummary {
	precioConIva := productoVendido.PrecioCompra + (productoVendido.PrecioCompra * (productoVendido.Iva / 100))
	beneficioProducto := productoVendido.PrecioFinal - precioConIva
	ivaProducto := productoVendido.PrecioCompra * (productoVendido.Iva / 100)

	ps.AddCantidadvendida(int32(productoVendido.CantidadVendida))
	ps.AddCosteTotalProducto(float64(productoVendido.CantidadVendida) * float64(productoVendido.PrecioCompra))
	ps.AddValorVenta(float64(productoVendido.CantidadVendida) * productoVendido.PrecioFinal)
	ps.AddBeneficio(float64(productoVendido.CantidadVendida) * beneficioProducto)
	ps.AddIVAPagado(float64(productoVendido.CantidadVendida) * ivaProducto)

	return ps
}

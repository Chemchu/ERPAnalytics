package types

type Venta struct {
	ID                      string            `json:"_id"`
	Productos               []ProductoVendido `json:"productos"`
	DineroEntregadoEfectivo float64           `json:"dineroEntregadoEfectivo"`
	DineroEntregadoTarjeta  float64           `json:"dineroEntregadoTarjeta"`
	PrecioVentaTotalSinDto  float64           `json:"precioVentaTotalSinDto"`
	PrecioVentaTotal        float64           `json:"precioVentaTotal"`
	Cambio                  float64           `json:"cambio"`
	Cliente                 Cliente           `json:"cliente"`
	VendidoPor              Empleado          `json:"vendidoPor"`
	ModificadoPor           Empleado          `json:"modificadoPor"`
	Tipo                    string            `json:"tipo"`
	DescuentoEfectivo       float64           `json:"descuentoEfectivo"`
	DescuentoPorcentaje     float64           `json:"descuentoPorcentaje"`
	Tpv                     string            `json:"tpv"`
	CreatedAt               int64             `json:"createdAt,string"`
	UpdatedAt               int64             `json:"updatedAt,string"`
}

type VentasPorHora struct {
	Hora                  string  `json:"hora"`
	BeneficioHora         float64 `json:"beneficioHora"`
	TotalVentaHora        float64 `json:"totalVentaHora"`
	TotalEfectivoHora     float64 `json:"totalEfectivoHora"`
	TotalTarjetaHora      float64 `json:"totalTarjetaHora"`
	ProductosVendidosHora int     `json:"productosVendidosHora"`
	DineroDescontadoHora  float64 `json:"dineroDescontadoHora"`
}

type Devolucion struct {
	ID                 string             `json:"_id"`
	ProductosDevueltos []ProductoDevuelto `json:"productosDevueltos"`
	DineroDevuelto     float32            `json:"dineroDevuelto"`
	VentaOriginal      Venta              `json:"ventaOriginal"`
	Tpv                string             `json:"tpv"`
	Cliente            Cliente            `json:"cliente"`
	Trabajador         Empleado           `json:"trabajador"`
	ModificadoPor      Empleado           `json:"modificadoPor"`
	CreatedAt          string             `json:"createdAt"`
	UpdatedAt          string             `json:"updatedAt"`
}

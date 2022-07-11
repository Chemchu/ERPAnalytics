package types

type APIResponse struct {
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
	Data       string `json:"data,omitempty"`
}

type APIData struct {
	Message    string `json:"message,omitempty"`
	Successful bool   `json:"successful,omitempty"`
}

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

type ProductoVendido struct {
	ID              string  `json:"_id"`
	Nombre          string  `json:"nombre"`
	Familia         string  `json:"familia"`
	Proveedor       string  `json:"proveedor,omitempty"`
	PrecioCompra    float64 `json:"precioCompra"`
	PrecioVenta     float64 `json:"precioVenta"`
	PrecioFinal     float64 `json:"precioFinal"`
	CantidadVendida int     `json:"cantidadVendida"`
	Dto             float64 `json:"dto"`
	Iva             float64 `json:"iva"`
	Margen          float64 `json:"margen"`
	Ean             string  `json:"ean"`
}

type ProductoMasVendido struct {
	ID              string `json:"_id"`
	Nombre          string `json:"nombre"`
	Ean             string `json:"ean"`
	Familia         string `json:"familia"`
	CantidadVendida int    `json:"cantidadVendida"`
}

type FamiliaMasVendida struct {
	Familia         string `json:"familia"`
	CantidadVendida int    `json:"cantidadVendida"`
}

type Cliente struct {
	ID     string `json:"_id"`
	Nombre string `json:"nombre"`
	Calle  string `json:"calle"`
	Cp     string `json:"cp"`
	Nif    string `json:"nif"`
}

type Empleado struct {
	ID             string `json:"_id"`
	Nombre         string `json:"nombre"`
	Apellidos      string `json:"apellidos"`
	Dni            string `json:"dni"`
	Email          string `json:"email"`
	FechaAlta      string `json:"fechaAlta"`
	Genero         string `json:"genero"`
	HorasPorSemana int    `json:"horasPorSemana"`
	Rol            string `json:"rol"`
}

type Cierre struct {
	ID                   string   `json:"_id"`
	Tpv                  string   `json:"tpv"`
	AbiertoPor           Empleado `json:"abiertoPor"`
	CerradoPor           Empleado `json:"cerradoPor"`
	Apertura             string   `json:"apertura"`
	Cierre               string   `json:"cierre"`
	CajaInicial          float64  `json:"cajaInicial"`
	NumVentas            int32    `json:"numVentas"`
	VentasEfectivo       float64  `json:"ventasEfectivo"`
	VentasTarjeta        float64  `json:"ventasTarjeta"`
	VentasTotales        float64  `json:"ventasTotales"`
	DineroEsperadoEnCaja float32  `json:"dineroEsperadoEnCaja"`
	DineroRealEnCaja     float32  `json:"dineroRealEnCaja"`
	DineroRetirado       float32  `json:"dineroRetirado"`
	FondoDeCaja          float32  `json:"fondoDeCaja"`
	Beneficio            float32  `json:"beneficio"`
	Nota                 string   `json:"nota"`
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

type ProductoDevuelto struct {
	ID               string  `json:"_id"`
	Nombre           string  `json:"nombre"`
	Familia          string  `json:"familia"`
	Proveedor        string  `json:"proveedor"`
	PrecioCompra     float32 `json:"precioCompra"`
	PrecioVenta      float32 `json:"precioVenta"`
	PrecioFinal      float32 `json:"precioFinal"`
	CantidadDevuelta int32   `json:"cantidadDevuelta"`
	Dto              float32 `json:"dto"`
	Iva              float32 `json:"iva"`
	Margen           float32 `json:"margen"`
	Ean              string  `json:"ean"`
}

type Producto struct {
	ID              string  `json:"_id"`
	Nombre          string  `json:"nombre"`
	Familia         string  `json:"familia"`
	Proveedor       string  `json:"proveedor"`
	PrecioCompra    float32 `json:"precioCompra"`
	PrecioVenta     float32 `json:"precioVenta"`
	Iva             float32 `json:"iva"`
	Margen          float32 `json:"margen"`
	Ean             string  `json:"ean"`
	Promociones     string  `json:"promociones"`
	Alta            bool    `json:"alta"`
	Cantidad        int32   `json:"cantidad"`
	CantidadRestock int32   `json:"cantidadRestock"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
}

type Summary struct {
	VentasPorHora             []VentasPorHora      `json:"ventasPorHora"`
	ProductosMasVendidos      []ProductoMasVendido `json:"productosMasVendidos"`
	FamiliasMasVendidas       []FamiliaMasVendida  `json:"familiasMasVendidas"`
	Beneficio                 float64              `json:"beneficio"`
	TotalVentas               float64              `json:"totalVentas"`
	TotalEfectivo             float64              `json:"totalEfectivo"`
	TotalTarjeta              float64              `json:"totalTarjeta"`
	NumVentas                 int                  `json:"numVentas"`
	MediaVentas               float64              `json:"mediaVentas"`
	MediaCantidadVenida       float64              `json:"mediaCantidadVenida"`
	CantidadProductosVendidos int                  `json:"cantidadProductosVendidos"`
	DineroDescontado          float64              `json:"dineroDescontado"`
	IVAPagado                 float64              `json:"ivaPagado"`
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

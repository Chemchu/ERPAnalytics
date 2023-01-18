package types

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

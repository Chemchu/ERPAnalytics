package types

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

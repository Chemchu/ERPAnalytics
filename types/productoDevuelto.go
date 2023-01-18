package types

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

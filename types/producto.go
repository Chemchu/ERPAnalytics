package types

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

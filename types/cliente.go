package types

type Cliente struct {
	ID     string `json:"_id"`
	Nombre string `json:"nombre"`
	Calle  string `json:"calle"`
	Cp     string `json:"cp"`
	Nif    string `json:"nif"`
}

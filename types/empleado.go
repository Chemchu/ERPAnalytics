package types

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

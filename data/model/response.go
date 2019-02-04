package model

import "time"

//RUsuario struct
type RUsuario struct {
	ID         int
	Nombre     string
	Correo     string
	Contrasena string
}

//RFoto struct
type RFoto struct {
	ID         int
	NombreFoto string
	Fecha      time.Time
}

//RComentario struct
type RComentario struct {
	ID    int
	Texto string
	Fecha time.Time
}

//RLogin struct
type RLogin struct {
	ID         string
	Correo     string
	Contrasena string
}

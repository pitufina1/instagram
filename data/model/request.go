package model

import "time"

//Ususario struct
type Usuario struct {
	Nombre     string
	Correo     string
	Contrasena string
}

//Filtro struct
type Filtro struct {
	Correo string
	Fecha  time.Time
}

//Foto struct
type Foto struct {
	NombreFoto string
	Fecha      time.Time
}

//Comentario struct
type Comentario struct {
	Texto string
	Fecha time.Time
}

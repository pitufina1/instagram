package dataclient

import (
	"database/sql"
	"fmt"
	"instagram/data/model"
	"time"

	_ "github.com/go-sql-driver/mysql" ///El driver se registra en database/sql en su función Init(). Es usado internamente por éste
)

//InsertarUsuario test
func InsertarUsuario(objeto *model.Usuario) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")

	if err != nil {
		panic(err.Error()) //si se abre bien
	}

	defer db.Close() //cerrar la conexion
	insert, err := db.Query("INSERT INTO Usuario(nombre, correo, contrasena) VALUES (?, ?, ?)", objeto.Nombre, objeto.Correo, objeto.Contrasena)
	// Inserta un nuevo usuario en la base de datos
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

//Listar registro de usuarios test
func ListarRegistrosUsuarios(objeto *model.Filtro) []model.RUsuario {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	comando := "SELECT * FROM Usuario WHERE (correo <= '" + objeto.Correo + "')"
	fmt.Println(comando)
	query, err := db.Query("SELECT * FROM Usuario WHERE (correo >= ?)", objeto.Correo)

	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	resultado := make([]model.RUsuario, 0)
	for query.Next() {
		var fila = model.RUsuario{}

		err = query.Scan(&fila.ID, &fila.Nombre, &fila.Correo, &fila.Contrasena)
		if err != nil {
			panic(err.Error())
		}
		resultado = append(resultado, fila)
	}
	return resultado
}

//Login Usuario test
func InsertarLogin(objeto *model.Login) string {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	//Consultamos todos los usuarios registrados
	comando := "SELECT Contrasena FROM Usuario WHERE (Correo ='" + objeto.Correo + "')"
	fmt.Println(comando)
	query, err := db.Query("SELECT Contrasena FROM Usuario WHERE (Correo = '" + objeto.Correo + "')")

	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	var resultado string
	for query.Next() {

		err = query.Scan(&resultado)
		if err != nil {
			panic(err.Error())
		}
	}
	return resultado
}

//InsertarFoto test
func InsertarFoto(objeto *model.Foto) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")

	if err != nil {
		panic(err.Error()) //si se abre bien
	}

	defer db.Close() //cerrar la conexion
	insert, err := db.Query("INSERT INTO Foto(ID, NombreFoto, Fecha) VALUES (?, ?, utc_timestamp())", objeto.NombreFoto)
	// Inserta una foto en la base de datos
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

//ListarFotos test
func ListarFotos(objeto *model.Filtro) []model.RFoto {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	comando := "SELECT ID FROM Foto WHERE (fecha <= '" + objeto.Fecha.Format(time.RFC3339) + "')"
	fmt.Println(comando)
	query, err := db.Query("SELECT ID FROM Foto WHERE (fecha >= ?)", objeto.Fecha.Format(time.RFC3339))

	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	resultado := make([]model.RFoto, 0)
	for query.Next() {
		var fila = model.RFoto{}

		err = query.Scan(&fila.ID, &fila.NombreFoto, &fila.Fecha)
		if err != nil {
			panic(err.Error())
		}
		resultado = append(resultado, fila)
	}
	return resultado
}

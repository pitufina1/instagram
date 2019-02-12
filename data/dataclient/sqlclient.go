package dataclient

import (
	"database/sql"
	"fmt"
	"instagram/data/model"

	_ "github.com/go-sql-driver/mysql" ///El driver se registra en database/sql en su función Init(). Es usado internamente por éste
)

//InsertarUsuario test
func InsertarUsuario(objeto *model.Usuario) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	insert, err := db.Query("INSERT INTO Usuario(nombre, correo, contrasena) VALUES (?, ?, ?)", objeto.Nombre, objeto.Correo, objeto.Contrasena)
	// Inserta un nuevo usuario en la base de datos
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

//Listar registro de usuarios test
func ListarRegistrosUsuarios(correo string) int {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	comando := "SELECT ID FROM Usuario WHERE (correo = '" + correo + "')"
	fmt.Println(comando)
	query, err := db.Query("SELECT ID FROM Usuario WHERE (correo = '" + correo + "')")

	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	var resultado int
	for query.Next() {

		err := query.Scan(&resultado)
		if err != nil {
			panic(err.Error())
		}
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
func InsertarFoto(nombrefoto string, id int) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close() //cerrar la conexion
	insert, err := db.Query("INSERT INTO Foto(NombreFoto, Usuario_ID) VALUES (?, ?)", nombrefoto, id)
	// Inserta una foto en la base de datos
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

//ListarFotos test
func ListarFotos() []model.RFoto {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	comando := "SELECT ID, NombreFoto FROM Foto"
	fmt.Println(comando)
	query, err := db.Query("SELECT ID, NombreFoto FROM Foto")

	if err != nil {
		panic(err.Error())
	}

	resultado := make([]model.RFoto, 0)
	for query.Next() {
		var foto = model.RFoto{}

		err = query.Scan(&foto.ID, &foto.NombreFoto)
		if err != nil {
			panic(err.Error())
		}
		resultado = append(resultado, foto)
	}
	return resultado
}

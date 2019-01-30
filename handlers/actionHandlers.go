package handlers

import (
	"encoding/json"
	"fmt"
	client "instagram/data/dataclient"
	"instagram/data/model"
	"io/ioutil"
	"net/http"
	"strings"
)

//Insert Función que inserta una foto en la base de datos local
func Insert(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathEnvioFoto {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)

	if e == nil {
		var foto model.Foto
		enTexto := string(bytes)
		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &foto)

		if foto.NombreFoto != "" {
			foto.NombreFoto = strings.ToUpper(foto.NombreFoto)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "La foto está vacía")
			return
		}

		w.WriteHeader(http.StatusOK)

		w.Header().Add("Content-Type", "application/json")

		respuesta, _ := json.Marshal(foto)
		fmt.Fprint(w, string(respuesta))

		go client.InsertarFoto(&foto)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}
}

//Insert Función que inserta un usuario en la base de datos local
func InsertUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathInsertarUsuario {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)

	if e == nil {
		var usuario model.Usuario
		enTexto := string(bytes)
		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &usuario)

		fmt.Println(usuario.Nombre)

		if usuario.Nombre == "" || usuario.Correo == "" || usuario.Contrasena == "" {
			usuario.Nombre = strings.ToUpper(usuario.Nombre)

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "El usuario está vacío")
			return
		}

		w.WriteHeader(http.StatusOK)

		w.Header().Add("Content-Type", "application/json")

		respuesta, _ := json.Marshal(usuario)
		fmt.Fprint(w, string(respuesta))

		go client.InsertarUsuario(&usuario)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}
}

//List Función que devuelve el listado de usuarios de la base de datos dado un filtro
func ListUsuarios(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathListadoUsuarios {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)

	if e == nil {
		var filtro model.Filtro
		e = json.Unmarshal(bytes, &filtro)

		if e == nil {
			listausuarios := client.ListarRegistrosUsuarios(&filtro)

			w.WriteHeader(http.StatusOK)

			w.Header().Add("Content-Type", "application/json")

			respuesta, _ := json.Marshal(&listausuarios)
			fmt.Fprint(w, string(respuesta))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "La petición no pudo ser parseada")
			fmt.Fprintln(w, e.Error())
			return
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}
}

//List Función que devuelve las fotos de la base de datos dado un filtro
func List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathListadoFotos {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)

	if e == nil {
		var filtro model.Filtro
		e = json.Unmarshal(bytes, &filtro)

		if e == nil {
			lista := client.ListarRegistros(&filtro)

			w.WriteHeader(http.StatusOK)

			w.Header().Add("Content-Type", "application/json")

			respuesta, _ := json.Marshal(&lista)
			fmt.Fprint(w, string(respuesta))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "La foto no pudo ser parseada")
			fmt.Fprintln(w, e.Error())
			return
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}
}

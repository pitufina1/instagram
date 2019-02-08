package handlers

import (
	"encoding/json"
	"fmt"
	client "instagram/data/dataclient"
	"instagram/data/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
)

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
			fmt.Fprintln(w, "El registro está vacío")
			return
		}

		//Encripta la contraseña
		hash, err := bcrypt.GenerateFromPassword([]byte(usuario.Contrasena), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
		}
		hashComoCadena := string(hash)
		usuario.Contrasena = hashComoCadena

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

//LoginUsuario Función que devuelve el usuario logueado de la base de datos dado un filtro
func LoginUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathLoginUsuario {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)
	resp := false

	if e == nil {
		var usuario model.Login
		enTexto := string(bytes)
		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &usuario)

		fmt.Println(usuario.Correo)

		if usuario.Correo == "" || usuario.Contrasena == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "El registro está vacío")
			return
		}

		//Contraseña de la Base de Datos
		contrasena := client.InsertarLogin(&usuario)

		//Comprueba que las dos contraseñas sean iguales
		if err := bcrypt.CompareHashAndPassword([]byte(contrasena), []byte(usuario.Contrasena)); err != nil {
			fmt.Printf("No estas logeado")
			//fmt.Println(usuario.Contrasena)
		} else {
			resp = true
			setSession(usuario.Correo, w)
			fmt.Printf("Usuario logeado")
			getID(r)

		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, resp)
	}
	fmt.Fprintln(w, resp)
}

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

var router = mux.NewRouter()

/*func loginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")
	redirectTarget := "/"
	if name != "" && pass != "" {
		// .. check credentials ..
		setSession(name, response)
		redirectTarget = "/main"
	}
	http.Redirect(response, request, redirectTarget, 302)
}*/

func logoutHandler(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}

func setSession(name string, response http.ResponseWriter) {
	value := map[string]string{
		"name": name,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func getID(request *http.Request) (name string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			name = cookieValue["name"]
		}
	}
	return name
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

// UploadFile sube el archivo al servidor
func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(2500)

	file, fileInfo, err := r.FormFile("archivo")

	f, err := os.OpenFile("./files/"+fileInfo.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	//aki deberían ir las variables que me asocien el archivo que quiero subir correspondiente a un ID_usuario

	if err != nil {
		log.Fatal(err)
		return
	}

	defer f.Close()

	io.Copy(f, file)

	fmt.Fprintf(w, fileInfo.Filename)
}

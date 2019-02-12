package handlers

import "net/http"

//PathInicio Ruta raíz
const PathInicio string = "/"

//PathJSFiles Ruta a la carpeta de scripts de javascript
const PathJSFiles string = "/js/"

//PathCSSFiles Ruta a la carpeta de estilos css
const PathCSSFiles string = "/css/"

//PathInsertarUsuario Ruta de insertar usuario
const PathInsertarUsuario string = "/insertarusuario"

//PathListadoUsuarios Ruta de obtención de las usuarios de hoy
const PathLoginUsuario string = "/loginusuario"

//PathLogin Ruta de pagina de login de usuario
const PathLogin string = "/login"

//PathInsertarFoto Ruta para
const PathInsertarFoto string = "/upload"

//PathInsertarFoto Ruta para
const PathUpload string = "/upload"

//PathMain Ruta raíz
const PathMain string = "/main"

//PathMain Ruta raíz
const PathListadoFotos string = "/lista"

//ManejadorHTTP encapsula como tipo la función de manejo de peticiones HTTP, para que sea posible almacenar sus referencias en un diccionario
type ManejadorHTTP = func(w http.ResponseWriter, r *http.Request)

//Lista es el diccionario general de las peticiones que son manejadas por nuestro servidor
var Manejadores map[string]ManejadorHTTP

func init() {
	Manejadores = make(map[string]ManejadorHTTP)
	Manejadores[PathInicio] = IndexFile
	Manejadores[PathMain] = MainFile
	Manejadores[PathJSFiles] = JsFile
	Manejadores[PathCSSFiles] = CssFile
	Manejadores[PathInsertarUsuario] = InsertUsuario
	Manejadores[PathLogin] = Login
	Manejadores[PathLoginUsuario] = LoginUsuario
	Manejadores[PathInsertarFoto] = Upload
	Manejadores[PathUpload] = Upload
	Manejadores[PathListadoFotos] = ListadoFotos
}

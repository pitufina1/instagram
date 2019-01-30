package handlers

import "net/http"

//PathInicio Ruta raíz
const PathInicio string = "/"

//PathMain Ruta raíz
const PathMain string = "/main"

//PathJSFiles Ruta a la carpeta de scripts de javascript
const PathJSFiles string = "/js/"

//PathCSSFiles Ruta a la carpeta de estilos css
const PathCSSFiles string = "/css/"

//PathInsertarUsuario Ruta de insertar usuario
const PathInsertarUsuario string = "/insertarusuario"

//PathListadoUsuarios Ruta de obtención de las usuarios de hoy
const PathListadoUsuarios string = "/listausuarios"

//PathEnvioFoto Ruta de envío de fotos
const PathEnvioFoto string = "/envio"

//PathListadoFotos Ruta de obtención de las fotos de hoy
const PathListadoFotos string = "/lista"

//PathLogin Ruta de logueo de usuario
const PathLogin string = "/login"

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
	Manejadores[PathEnvioFoto] = Insert
	Manejadores[PathInsertarUsuario] = InsertUsuario
	Manejadores[PathListadoUsuarios] = List
	Manejadores[PathListadoFotos] = List
	Manejadores[PathLogin] = Login

}

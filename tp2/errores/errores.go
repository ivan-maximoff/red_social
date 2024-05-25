package errores

type ErrorLeerArchivo struct{}

func (e ErrorLeerArchivo) Error() string {
	return "ERROR: Lectura de archivos"
}

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "ERROR: Comando invalido"
}

type ErrorExcesoParametros struct{}

func (e ErrorExcesoParametros) Error() string {
	return "ERROR: Cantidad Invalida de Argumentos"
}

type UsuarioLoggeado struct{}

func (e UsuarioLoggeado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type UsuarioInexistente struct{}

func (e UsuarioInexistente) Error() string {
	return "Error: usuario no existente"
}

type UsuarioNoLoggeado struct{}

func (e UsuarioNoLoggeado) Error() string {
	return "Error: no habia usuario loggeado"
}

type ErrorProximoFeed struct{}

func (e ErrorProximoFeed) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver"
}

type ErrorLikearPost struct{}

func (e ErrorLikearPost) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}

type ErrorMostrarLikes struct{}

func (e ErrorMostrarLikes) Error() string {
	return "Error: Post inexistente o sin likes"
}

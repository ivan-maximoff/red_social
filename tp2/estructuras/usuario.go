package aplicacion

// Usuario modela a un usuario de algogram
type Usuario interface {

	//Devuelve el nombre de Usuario
	NombreUsuario() string

	//Devuelve el ID de Usuario
	IdUsuario() int

	//Recibe ID y distancia al Post que guarda en sus Proximos Posts para ver
	AgregarPostFeed(a, b int)

	//Devuelve el ID del Post Actual
	PostActual() int

	//Avanza al siguiente Post del Feed
	SiguientePost()

	//Devuelve true si ya no quedan publicaciones para ver
	SinPublicaciones() bool
}

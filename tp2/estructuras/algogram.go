package aplicacion

type Algogram interface {

	//Logea un Usuario
	Login(string)

	//Deslogea el usuario actual
	Logout()

	//True si hay un Logeado, False si no lo hay
	HayLogeado() bool

	//Publica un Post con el mensaje dado
	Publicar(string)

	//Avanza a la siguiente publicacion en el feed del Logeado
	Proximo()

	//Likea el post del ID correspondiente
	Likear(int)

	//Muestra que usuarios le dieron like a la publicacion solicitada
	Mostrar(int)

	//True si el usuario existe, False si no existe
	ExisteUsuario(string) bool

	//Devuelve el usuario logeado
	UsuarioActual() Usuario

	//True si Existe el Post dado
	ExistePost(int) bool

	//Devuelve el Post pedido
	PostPedido(int) Post
}

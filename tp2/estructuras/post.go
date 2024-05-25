package aplicacion

// Post modela un post/publicacion de algogram
type Post interface {
	//Le da like al post
	Likear(string)

	//Muestra los usarios que le dieron like a la publicacion en orden alfabetico
	MostrarLikes()

	//Devuelve la cantidad de like del post
	Likes() int

	//Imprime en pantalla el ID, autor, mensaje y cantidad de likes del post
	VerPost()
}

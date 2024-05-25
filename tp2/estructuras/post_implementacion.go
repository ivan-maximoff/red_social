package aplicacion

import (
	"fmt"
	"strings"
	TDADiccionario "tdas/diccionario"
)

// STRUCTS
type post struct {
	IDPost     int
	autor      *Usuario
	mensaje    string
	likeadoPor TDADiccionario.DiccionarioOrdenado[string, string]
}

//FUNCIONES PRINCIPALES

// Crea y devuelve un post
func CrearPost(mensaje string, autor *Usuario, idPost int) Post {
	post := new(post)
	post.IDPost = idPost
	post.likeadoPor = TDADiccionario.CrearABB[string, string](func(i1, i2 string) int { return strings.Compare(i1, i2) })
	post.mensaje = mensaje
	post.autor = autor
	return post
}

func (post *post) Likear(nombre string) {
	if !post.likeadoPor.Pertenece(nombre) {
		post.likeadoPor.Guardar(nombre, nombre)
	}
	fmt.Println("Post likeado")
}

func (post post) Likes() int {
	return post.likeadoPor.Cantidad()
}

func (post post) MostrarLikes() {
	fmt.Println("El post tiene", post.Likes(), "likes:")
	post.likeadoPor.Iterar(func(clave, valor string) bool {
		fmt.Printf("\t%s\n", clave)
		return true
	})
}

func (post post) VerPost() {
	fmt.Println("Post ID", post.IDPost)
	fmt.Println((*post.autor).NombreUsuario(), "dijo:", post.mensaje)
	fmt.Println("Likes:", post.likeadoPor.Cantidad())
}

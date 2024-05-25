package aplicacion

import (
	errores "algogram/errores"
	"bufio"
	"fmt"
	"os"

	TDADiccionario "tdas/diccionario"
)

type algogram struct {
	hashUsuarios TDADiccionario.Diccionario[string, Usuario]
	hashPosts    TDADiccionario.Diccionario[int, Post]
	logeado      *Usuario
}

// lee el archivo de usuarios y crea un hash.
func inicializarUsuarios(usuarios string) TDADiccionario.Diccionario[string, Usuario] {
	archivo, err := os.Open(usuarios)
	if err != nil {
		fmt.Println(errores.ErrorLeerArchivo{})
		return nil
	}
	defer archivo.Close()
	hash := TDADiccionario.CrearHash[string, Usuario]()
	s := bufio.NewScanner(archivo)
	var id int
	for s.Scan() {
		nombre := s.Text()
		usuario := CrearUsuario(id, nombre)
		hash.Guardar(nombre, usuario)
		id++
	}
	err = s.Err()
	if err != nil {
		fmt.Println(errores.ErrorLeerArchivo{})
	}
	return hash
}

// Crea y Devuelve un Algogram
func CrearAlgogram(archivo string) Algogram {
	algogram := new(algogram)
	algogram.hashUsuarios = inicializarUsuarios(archivo)
	algogram.hashPosts = TDADiccionario.CrearHash[int, Post]()
	algogram.logeado = nil
	return algogram

}

func (algo *algogram) Login(nombre string) {
	usuario := algo.hashUsuarios.Obtener(nombre)
	algo.logeado = &usuario
	mensaje := "Hola " + nombre
	fmt.Println(mensaje)
}

func (algo *algogram) Logout() {
	algo.logeado = nil
	fmt.Println("Adios")
}

func (algo algogram) HayLogeado() bool {
	return algo.logeado != nil
}

func (algo algogram) ExisteUsuario(usuario string) bool {
	return algo.hashUsuarios.Pertenece(usuario)
}

func (algo algogram) UsuarioActual() Usuario {
	if !algo.HayLogeado() {
		panic("No hay usuario Logeado")
	}
	return (*algo.logeado)
}

func (algo *algogram) Publicar(mensaje string) {
	id, idUsuAct, nuevoPost := algo.inicializarNuevoPost(mensaje)
	algo.hashPosts.Guardar(id, nuevoPost)
	algo.agregarPostenUsuarios(idUsuAct, id)
	fmt.Println("Post publicado")
}

func (algo algogram) ExistePost(id int) bool {
	return algo.hashPosts.Pertenece(id)
}

func (algo algogram) PostPedido(id int) Post {
	return algo.hashPosts.Obtener(id)
}

func (algo *algogram) Proximo() {
	post := algo.hashPosts.Obtener((*algo.logeado).PostActual())
	post.VerPost()
	(*algo.logeado).SiguientePost()
}

func (algo *algogram) Likear(id int) {
	nombreUsuario := (*algo.logeado).NombreUsuario()
	algo.hashPosts.Obtener(id).Likear(nombreUsuario)
}

func (algo *algogram) Mostrar(id int) {
	algo.hashPosts.Obtener(id).MostrarLikes()
}

//Herramientas Auxiliares:

// Crea un nuevo Post con la informacion correspondiente
func (algo *algogram) inicializarNuevoPost(mensaje string) (int, int, Post) {
	usuarioActual := algo.logeado
	idUsuarioActual := (*algo.logeado).IdUsuario()
	id := algo.hashPosts.Cantidad()
	nuevoPost := CrearPost(mensaje, usuarioActual, id)
	return id, idUsuarioActual, nuevoPost
}

// Agrega el post al feed de cada Usuario
func (algogram *algogram) agregarPostenUsuarios(idUsuAct, idPost int) {
	iter := algogram.hashUsuarios.Iterador()
	for iter.HaySiguiente() {
		_, dato := iter.VerActual()
		distancia := calcularDistancia(idUsuAct, dato.IdUsuario())
		if distancia != 0 {
			dato.AgregarPostFeed(idPost, distancia)
		}
		iter.Siguiente()
	}
}

// calcula la diferencia entre dos numeros. Siempre positiva
func calcularDistancia(actual, iterado int) int {
	distancia := actual - iterado
	return valorAbsoluto(distancia)
}

// calcula y devuelve el valor absoluto de un entero
func valorAbsoluto(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

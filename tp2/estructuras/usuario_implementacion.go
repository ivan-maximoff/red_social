package aplicacion

import (
	TDAHeap "tdas/heap"
)

// STRUCTS
type prioridad struct {
	IDpost    int
	distancia int
}

type usuario struct {
	nombre       string
	id           int
	proximosPost TDAHeap.ColaPrioridad[prioridad]
}

//FUNCIONES PRINCIPALES

// Crea y devuelve un Usuario
func CrearUsuario(id int, nombre string) Usuario {
	usuario := new(usuario)
	usuario.nombre = nombre
	usuario.id = id
	usuario.proximosPost = TDAHeap.CrearHeap[prioridad](comparar)
	return usuario
}

func (usu *usuario) IdUsuario() int {
	return usu.id
}

func (usu *usuario) NombreUsuario() string {
	return usu.nombre
}

func (usu usuario) SinPublicaciones() bool {
	return usu.proximosPost.EstaVacia()
}

func (usu *usuario) AgregarPostFeed(id, distancia int) {
	prioridad := crearPrioridad(id, distancia)
	usu.proximosPost.Encolar(prioridad)
}

func (usu *usuario) PostActual() int {
	usu.panicVacio()
	return usu.proximosPost.VerMax().IDpost
}

func (usu *usuario) SiguientePost() {
	usu.panicVacio()
	usu.proximosPost.Desencolar()
}

// Auxiliares
func crearPrioridad(id, distancia int) prioridad {
	prioridad := prioridad{}
	prioridad.IDpost = id
	prioridad.distancia = distancia

	return prioridad
}
func (usu usuario) panicVacio() {
	if usu.SinPublicaciones() {
		panic("No hay publicaciones")
	}
}

func comparar(p1, p2 prioridad) int {
	dif := p2.distancia - p1.distancia
	if dif == 0 {
		return p2.IDpost - p1.IDpost
	}
	return dif
}

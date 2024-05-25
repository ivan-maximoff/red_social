package main

import (
	"algogram/errores"
	estructuras "algogram/estructuras"
	"strconv"
	"strings"
)

const (
	CANT_PARAMETROS_LOGIN   = 1
	CANT_PARAMETROS_LOGOUT  = 0
	CANT_MIN_MENSAJE        = 0
	CANT_PARAMETROS_PROXIMO = 0
	CANT_PARAMETROS_LIKEAR  = 1
	CANT_PARAMETROS_MOSTRAR = 1
)

// Verifica que la cantidad de argumentos coincida con la esperada, en caso contrario
// devuelve el error
func validarArgumentos(cantidad_argumentos, argumentos_esperados int) error {
	if cantidad_argumentos != argumentos_esperados {
		return errores.ErrorExcesoParametros{}
	}
	return nil
}

// Verifica que el parametro para Login sea correcto, devuelve el nombre del usuario leido
// en caso contrario devuelve el error
func valdiacionLogin(algogram estructuras.Algogram, parametro []string) (error, string) {
	if algogram.HayLogeado() {
		return errores.UsuarioLoggeado{}, ""
	}
	nombre := strings.Join(parametro, " ")
	if !algogram.ExisteUsuario(nombre) {
		return errores.UsuarioInexistente{}, ""
	}
	return nil, nombre
}

// Verifica que el parametro para Logout sea correcto
// en caso contrario devuelve el error
func validarLogout(algogram estructuras.Algogram, parametro []string) error {
	err := validarArgumentos(len(parametro), CANT_PARAMETROS_LOGOUT)
	if err != nil {
		return err
	}
	if !algogram.HayLogeado() {
		return errores.UsuarioNoLoggeado{}
	}
	return nil
}

// Verifica que el parametro para Publicar sea correcto
// en caso contrario devuelve el error
func validarPublicar(algogram estructuras.Algogram) error {
	if !algogram.HayLogeado() {
		return errores.UsuarioNoLoggeado{}
	}
	return nil
}

// Verifica que el parametro para Proximo sea correcto, devuelve el nombre del usuario y el id del post
// en caso contrario devuelve el error
func validarProximo(algogram estructuras.Algogram, parametro []string) error {
	err := validarArgumentos(len(parametro), CANT_PARAMETROS_PROXIMO)
	if err != nil || !algogram.HayLogeado() {
		return errores.ErrorProximoFeed{}
	}
	usuario := algogram.UsuarioActual()
	if usuario.SinPublicaciones() {
		return errores.ErrorProximoFeed{}
	}
	return nil
}

// Verifica que el parametro para Likear sea correcto, devuelve el id del post leído
// en caso contrario devuelve el error
func validarLikear(algogram estructuras.Algogram, parametro []string) (error, int) {
	idInt, err1 := strconv.Atoi(parametro[0])
	if err1 != nil {
		return err1, 0
	}
	id := idInt
	err2 := validarArgumentos(len(parametro), CANT_PARAMETROS_LIKEAR)
	if err2 != nil || !algogram.HayLogeado() || !algogram.ExistePost(id) {
		return errores.ErrorLikearPost{}, 0
	}
	return nil, id
}

// Verifica que el parametro para Mostrar sea correcto, devuelve el ID del post leído
// en caso contrario devuelve el error
func validarMostrar(algogram estructuras.Algogram, parametro []string) (error, int) {
	err := validarArgumentos(len(parametro), CANT_PARAMETROS_MOSTRAR)
	if err != nil {
		return err, 0
	}
	idInt, err1 := strconv.Atoi(parametro[0])
	if err1 != nil {
		return err1, 0
	}
	id := idInt
	if !algogram.ExistePost(id) || algogram.PostPedido(id).Likes() == 0 {
		return errores.ErrorMostrarLikes{}, 0
	}
	return nil, id
}

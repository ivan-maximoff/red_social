package main

import (
	"algogram/errores"
	estructuras "algogram/estructuras"
	"fmt"
	"strings"
	TDADiccionario "tdas/diccionario"
)

const (
	LOGIN         = "login"
	LOGOUT        = "logout"
	PUBLICAR_POST = "publicar"
	VER_PROXIMO   = "ver_siguiente_feed"
	LIKEAR_POST   = "likear_post"
	MOSTRAR_LIKES = "mostrar_likes"
)

var dicFunciones = TDADiccionario.CrearHash[string, func(estructuras.Algogram, []string)]()

// Asignar las funciones al diccionario global una sola vez cuando se ejecuta el programa
func init() {
	dicFunciones.Guardar(LOGIN, login)
	dicFunciones.Guardar(LOGOUT, logout)
	dicFunciones.Guardar(PUBLICAR_POST, publicar)
	dicFunciones.Guardar(VER_PROXIMO, proximo)
	dicFunciones.Guardar(LIKEAR_POST, likear)
	dicFunciones.Guardar(MOSTRAR_LIKES, mostrar)
}

// Ejecuta el comando correspondiente
func EjecutarComandos(algogram estructuras.Algogram, comando string, parametro []string) {
	if comando == "" {
		fmt.Println(errores.ErrorParametros{})
		return
	}
	if dicFunciones.Pertenece(comando) {
		funcion := dicFunciones.Obtener(comando)
		funcion(algogram, parametro)
	} else {
		fmt.Println(errores.ErrorParametros{})
	}
}

// Valida y Ejecuta Login
func login(algogram estructuras.Algogram, parametro []string) {
	err, nombre := valdiacionLogin(algogram, parametro)
	if err != nil {
		fmt.Println(err)
		return
	}
	algogram.Login(nombre)
}

// Valida y Ejecuta Logout
func logout(algogram estructuras.Algogram, parametro []string) {
	err := validarLogout(algogram, parametro)
	if err != nil {
		fmt.Println(err)
		return
	}
	algogram.Logout()
}

// Valida y Ejecuta Publicar
func publicar(algogram estructuras.Algogram, parametro []string) {
	err := validarPublicar(algogram)
	if err != nil {
		fmt.Println(err)
		return
	}
	mensaje := strings.Join(parametro, " ")
	algogram.Publicar(mensaje)
}

// Valida y Ejecuta Proximo
func proximo(algogram estructuras.Algogram, parametro []string) {
	err := validarProximo(algogram, parametro)
	if err != nil {
		fmt.Println(err)
		return
	}
	algogram.Proximo()
}

// Valida y Ejecuta Likear
func likear(algogram estructuras.Algogram, parametro []string) {
	err, id := validarLikear(algogram, parametro)
	if err != nil {
		fmt.Println(err)
		return
	}
	algogram.Likear(id)
}

// Valida y Ejecuta Mostrar
func mostrar(algogram estructuras.Algogram, parametro []string) {
	err, id := validarMostrar(algogram, parametro)
	if err != nil {
		fmt.Println(err)
		return
	}
	algogram.Mostrar(id)
}

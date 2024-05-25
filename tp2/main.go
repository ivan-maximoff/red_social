package main

import (
	"algogram/errores"
	estructuras "algogram/estructuras"
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	PRIMER_ARGUMENTO  = 1
	CANT_ARCHIVOS     = 1
	POSICION_USUARIOS = 0
)

func main() {
	parametros := os.Args[PRIMER_ARGUMENTO:] // ./algogram ARCHIVO_USUARIOS
	if len(parametros) != CANT_ARCHIVOS {
		fmt.Println(errores.ErrorParametros{})
		return
	}
	archivoUsuarios := parametros[POSICION_USUARIOS]
	algogram := estructuras.CrearAlgogram(archivoUsuarios)
	comandos(algogram)
}

func devolverComando(palabras []string) (string, []string) {
	var primerPalabra string
	var resto []string
	if len(palabras) >= 1 {
		primerPalabra = palabras[0]
	}
	if len(palabras) >= 2 {
		resto = palabras[1:]
	}
	return primerPalabra, resto
}

func comandos(algogram estructuras.Algogram) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linea := scanner.Text()
		lista := strings.Fields(linea)
		comando, parametro := devolverComando(lista)
		EjecutarComandos(algogram, comando, parametro)
	}
}

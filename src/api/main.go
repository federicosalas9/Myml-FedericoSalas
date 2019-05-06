package main

import (
	ginGonic "github.com/gin-gonic/gin"//si hay dos dependencias que se llamen igual agrego otro nombre adelante
	"github.com/mercadolibre/myml/src/api/controllers/ping"
	"github.com/mercadolibre/myml/src/api/controllers/myml"
)

const (
	port = ":8080" //no cambia a lo largo de la ejecucion del programa
)

var (
	router = ginGonic.Default() //se puede cambiar el valor pero no el tipo de dato
)

func main() {

	router.GET(
		"/ping", ping.Ping) //el parametro debe ser una funcion que no devuelva nada
	router.GET(
		"/myml/:userID", myml.GetInfoC)
	router.Run(port)
}

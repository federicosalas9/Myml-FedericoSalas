package main

import (
	ginGonic "github.com/gin-gonic/gin" //si hay dos dependencias que se llamen igual agrego otro nombre adelante
	"github.com/mercadolibre/Myml - Federico Salas/src/api/controllers/myml"
	"github.com/mercadolibre/Myml - Federico Salas/src/api/controllers/ping"
)

const (
	port = ":8080" //no cambia a lo largo de la ejecucion del programa
)

var ()

func main() {
	r := Start()
	r.Run(port)
}

func Start() *ginGonic.Engine {
	r := ginGonic.Default() //se puede cambiar el valor pero no el tipo de dato
	r.GET(
		"/ping", ping.Ping) //el parametro debe ser una funcion que no devuelva nada
	r.GET(
		"/myml/:userID", myml.GetInfoC)
	return r
}

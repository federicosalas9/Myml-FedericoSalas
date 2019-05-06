package ping

import ginGonic "github.com/gin-gonic/gin" //si hay dos dependencias que se llamen igual agrego otro nombre adelante

func Ping(context *ginGonic.Context) {
	context.String(200,"pong")
}
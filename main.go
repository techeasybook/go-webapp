package main

import (
	"gowebapp/handler"
	"gowebapp/router"
)

var (
	httpRouter = router.NewMuxRouter()
)

func main() {
	httpRouter.GET("/", handler.HealthCheck)
	httpRouter.GET("/getUser", handler.GetUser)
	httpRouter.SERVE("80")
}

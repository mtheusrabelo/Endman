package routes

import "github.com/gin-gonic/gin"

var r *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	r = gin.Default()
}

//Listen initializes the router
func Listen(port string) {
	go r.Run(port)
}

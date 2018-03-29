package main

import (
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	gin.SetMode(gin.ReleaseMode)
}
func main() {
	r := gin.Default()

	r.Use(gin.Recovery())
	r.Use(static.Serve("/", static.LocalFile("./public", true)))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	log.Println("http://localhost:8080/")
	r.Run(":8080")
}

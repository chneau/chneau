package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	gin.SetMode(gin.ReleaseMode)
	if runtime.GOOS == "windows" {
		gin.DisableConsoleColor()
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		log.Println("KEKK")
		os.Exit(0)
	}()
}

func main() {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./public", true)))

	log.Println("running on http://localhost:8080")
	r.Run(":8080")
}

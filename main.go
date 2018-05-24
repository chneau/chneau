package main

import (
	"log"
	"net/http"
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
}

func gracefulExit(server *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := server.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
		os.Exit(0)
	}()
}

func runServer(server *http.Server) {
	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpect")
		}
	}
}

func main() {

	r := gin.Default()

	r.Use(gin.Recovery())
	r.Use(static.Serve("/", static.LocalFile("./public", true)))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("http://localhost:8080")
	gracefulExit(server)
	runServer(server)
}

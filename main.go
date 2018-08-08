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
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		println()
		os.Exit(0)
	}()
}

func main() {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./public", true)))

	bashrc := "https://raw.githubusercontent.com/chneau/usefulCommands/master/.bashrc"
	r.GET("/.bashrc", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, bashrc)
	})
	r.GET("/bashrc", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, bashrc)
	})
	r.GET("/bash", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, bashrc)
	})
	r.GET("/b", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, bashrc)
	})

	log.Println("running on http://localhost:8080")
	r.Run(":8080")
}

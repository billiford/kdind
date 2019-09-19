package main

import (
	"flag"
	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
	port   = flag.String("p", "8080", "listen port")
)

func main() {
	flag.Parse()

	router = gin.Default()
	router.GET("/kill", func(c *gin.Context) {
		cmd := exec.Command("pkill", "dockerd")
		cmd.Run()
	})

	router.Run(fmt.Sprintf(":%s", *port))
}

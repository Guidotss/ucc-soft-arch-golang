package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Guidotss/ucc-soft-arch-golang.git/config"
	"github.com/gin-gonic/gin"
)

func main() {
	envs := config.LoadEnvs(".env")
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start server
	startServer(router, envs)
}

func startServer(router http.Handler, envs config.Envs) {
	serverPort := envs.Get("PORT")

	server := &http.Server{
		Addr:           ":" + serverPort,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		_ = fmt.Errorf("Error starting server: %v", err)
		panic(err)
	}
}

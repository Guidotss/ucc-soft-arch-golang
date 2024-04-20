package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Guidotss/ucc-soft-arch-golang.git/app/routes"
	"github.com/Guidotss/ucc-soft-arch-golang.git/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Cargar variables de entorno
	envs := config.LoadEnvs(".env")

	// Crear un nuevo router
	router := gin.Default()

	// Llamar a la función que define las rutas de la aplicación
	routes.AppRoutes(router)

	// Iniciar el servidor
	startServer(router, envs)
}

func startServer(router *gin.Engine, envs config.Envs) {
	serverPort := envs.Get("PORT")
	router.Use(cors.Default())

	server := &http.Server{
		Addr:           ":" + serverPort,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		_ = fmt.Errorf("error starting server: %v", err)
		panic(err)
	}
}

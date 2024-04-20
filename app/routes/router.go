package routes

import (
	"github.com/gin-gonic/gin"
)

// AppRoutes es la función que se encarga de definir las rutas de la aplicación
func AppRoutes(engine *gin.Engine /* db *gorm.DB */) {

	v1Routes := engine.Group("/v1")

	{
		AuthRoutes(v1Routes) // <- Ademas deberia recibir el controlador de autenticación
		HealthRoutes(v1Routes)
	}

}

package routes

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/adapter"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AppRoutes es la función que se encarga de definir las rutas de la aplicación
func AppRoutes(engine *gin.Engine, db *gorm.DB) {

	CoursesRoutes(engine, adapter.CourseAdapter(db))
	CategoriesRoutes(engine, adapter.CategoryAdapter(db))
	UsersRoutes(engine, adapter.UserAdapter(db))
	AuthRoutes(engine, adapter.AuthAdapter(db))
	InscriptionsRoutes(engine, adapter.InscriptionsAdapter(db))
	RatingRoutes(engine, adapter.RatingAdapter(db))
}

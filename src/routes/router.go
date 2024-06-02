package routes

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/adapter"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AppRoutes es la función que se encarga de definir las rutas de la aplicación
func AppRoutes(engine *gin.Engine, db *gorm.DB) {
	InscriptionController, InscriptionService := adapter.InscriptionsAdapter(db)
	UserController, UserService := adapter.UserAdapter(db)

	CoursesRoutes(engine, adapter.CourseAdapter(db))
	CategoriesRoutes(engine, adapter.CategoryAdapter(db))
	UsersRoutes(engine, UserController, UserService)
	AuthRoutes(engine, adapter.AuthAdapter(db))
	InscriptionsRoutes(engine, InscriptionController, InscriptionService)
	RatingRoutes(engine, adapter.RatingAdapter(db))
	CommentsRoutes(engine, adapter.CommentAdapter(db))

	engine.NoRoute(func(c *gin.Context) {
		c.Error(errors.NewError("NOT_FOUND", "Route not found", 404))
	})
}

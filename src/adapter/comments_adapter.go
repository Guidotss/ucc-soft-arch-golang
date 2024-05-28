package adapter

import (
	client "github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/comments"
	controllers "github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/comments"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"gorm.io/gorm"
)

func CommentAdapter(db *gorm.DB) *controllers.CommentsController {
	client := client.NewCommentsClient(db)
	service := services.NewCommentsService(client)
	return controllers.NewCommentsController(service)
}

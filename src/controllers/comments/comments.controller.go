package comments

import (
	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/comments"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"github.com/gin-gonic/gin"
)

type CommentsController struct {
	CommentsService services.ICommentsService
}

func NewCommentsController(service services.ICommentsService) *CommentsController {
	return &CommentsController{CommentsService: service}
}

func (c *CommentsController) NewComment(g *gin.Context) {
	var commentDto dto.CommentRequestResponseDto
	if err := g.BindJSON(&commentDto); err != nil {
		g.JSON(400, gin.H{
			"message": "La estructura del json es incorrecta",
		})
		return
	}

	response := c.CommentsService.NewComment(commentDto)
	g.JSON(201, gin.H{
		"response": response,
		"message":  "La comentario se registro con exito",
	})
}

package comments

import (
	"net/http"

	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/comments"
	customError "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/errors"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		g.Error(err)
		return
	}

	response, err := c.CommentsService.NewComment(commentDto)
	if err != nil {
		g.Error(err)
		return
	}
	g.JSON(201, gin.H{
		"response": response,
		"message":  "La comentario se registro con exito",
	})
}

func (c *CommentsController) GetCourseComments(g *gin.Context) {
	id := g.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		g.Error(customError.NewError("INVALID_UUID", "Invalid UUID", http.StatusBadRequest))
		return
	}
	response, err := c.CommentsService.GetCourseComments(uuid)
	if err != nil {
		g.Error(err)
		return
	}

	g.JSON(200, response)
}
func (c *CommentsController) UpdateComment(g *gin.Context) {
	var commentDto dto.CommentRequestResponseDto
	if err := g.ShouldBindJSON(&commentDto); err != nil {
		g.Error(customError.NewError("INVALID_INPUTS", "Invalid fields", http.StatusBadRequest))
		return
	}
	response, err := c.CommentsService.UpdateComment(commentDto)
	if err != nil {
		g.Error(err)
		return
	}
	g.JSON(200, gin.H{
		"response": response,
		"message":  "El comentario se actualizo con exito",
	})
}

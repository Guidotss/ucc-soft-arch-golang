package services

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/inscriptos"
	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/inscription"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
)

type IInscriptionService interface {
	Enroll(dto.EnrollRequestResponseDto) dto.EnrollRequestResponseDto
}

type inscriptionService struct {
	client inscriptos.InscriptosClient
}

func NewInscriptionService(client *inscriptos.InscriptosClient) IInscriptionService {
	return &inscriptionService{client: *client}
}

func (c *inscriptionService) Enroll(data dto.EnrollRequestResponseDto) dto.EnrollRequestResponseDto {
	var newEnroll = model.Inscriptos{
		CourseId: data.CourseId,
		UserId:   data.UserId,
	}
	enroll := c.client.Enroll(newEnroll)

	return dto.EnrollRequestResponseDto{
		CourseId: enroll.CourseId,
		UserId:   enroll.UserId,
	}
}

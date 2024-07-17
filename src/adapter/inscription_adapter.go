package adapter

import (
	client "github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/inscriptos"
	controllers "github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/inscriptions"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"gorm.io/gorm"
)

func InscriptionsAdapter(db *gorm.DB) (*controllers.InscriptionController, services.IInscriptionService) {
	client := client.NewInscriptionClient(db)
	service := services.NewInscriptionService(client)
	return controllers.NewInscriptionController(service), service
}

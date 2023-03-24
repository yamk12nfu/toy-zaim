package services

import "github.com/yamk12nfu/toy-zaim/app/domain/entities"

type ZaimService struct {
	ZaimHandler
}

func (service *ZaimService) GetZaimData() (entities.ZaimData, error) {
	return service.ZaimHandler.GetZaimData()
}

package usecases

import "github.com/yamk12nfu/toy-zaim/app/domain/entities"

type ZaimService interface {
	GetZaimData() (entities.ZaimData, error)
}

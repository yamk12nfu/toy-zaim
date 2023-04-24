package services

import "github.com/yamk12nfu/toy-zaim/app/domain/entities"

type ZaimHandler interface {
	GetMoneyData() ([]int64, error)
	GetZaimData() (entities.ZaimData, error)
}

package handler

import (
	"github.com/yamk12nfu/toy-zaim/app/interfaces/services"
	"github.com/yamk12nfu/toy-zaim/app/usecases/interactor"
)

type ZaimController struct {
	Interactor interactor.ZaimInteractor
}

func NewZaimController(zaimHandler services.ZaimHandler) *ZaimController {
	return &ZaimController{
		Interactor: interactor.ZaimInteractor{
			ZaimService: &services.ZaimService{
				ZaimHandler: zaimHandler,
			},
		},
	}
}

func (c *ZaimController)Sum() (int, error) {
	return c.Interactor.Sum()
}

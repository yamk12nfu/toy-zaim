package entities

import (
	gozaim "github.com/s-sasaki-0529/go-zaim"
)

type ZaimData struct {
	money      []gozaim.Money
	categories []gozaim.Category
	genres     []gozaim.Genre
	accounts   []gozaim.Account
}

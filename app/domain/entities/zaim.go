package entities

import (
	gozaim "github.com/s-sasaki-0529/go-zaim"
)

type ZaimData struct {
	Money      []gozaim.Money
	Categories []gozaim.Category
	Genres     []gozaim.Genre
	Accounts   []gozaim.Account
}

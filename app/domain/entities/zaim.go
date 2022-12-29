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

type MoneyJP struct {
	Date         string
	Mode         string
	Category     string
	Genre        string
	// From         string
	// To           string
	// Name         string
	// Comment      string
	// Place        string
	// CurrencyCode string
	// Income       int
	 Payment      int
	// Transfer     int
}

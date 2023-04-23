package interactor

import (
	gozaim "github.com/s-sasaki-0529/go-zaim"
	"github.com/yamk12nfu/toy-zaim/app/domain/entities"
	"github.com/yamk12nfu/toy-zaim/app/usecases/port"
)

type ZaimInteractor struct {
	ZaimService port.ZaimService
}

func (i *ZaimInteractor) Sum() (int, error) {
	data, err := i.ZaimService.GetZaimData()
	if err != nil {
		return -1, err
	}

	_, s := ConvertData(data)
	return s, nil
}

func ConvertData(data entities.ZaimData) ([]entities.MoneyJP, int) {
	sum := 0
	var money []entities.MoneyJP
	for _, v := range data.Money {
		pay, _, _ := GetAmount(v.Mode, v.Amount)
		cm := entities.MoneyJP{
			Date:     v.Date,
			Mode:     ConvertMode(v.Mode),
			Category: GetCategoryName(v.CategoryID, data.Categories),
			Genre:    GetGenreName(v.GenreID, data.Genres),
			Payment:  pay,
		}
		sum += pay

		money = append(money, cm)
	}

	return money, sum
}

func GetCategoryName(categoryID int, categories []gozaim.Category) string {
	for _, v := range categories {
		if v.ID == categoryID {
			return v.Name
		}
	}
	return ""
}

func GetGenreName(genreID int, genres []gozaim.Genre) string {
	for _, v := range genres {
		if v.ID == genreID {
			return v.Name
		}
	}
	return ""
}

func GetAmount(mode string, amount int) (int, int, int) {
	var p, i, t int

	switch mode {
	case "payment":
		p = amount
	case "income":
		i = amount
	case "transfer":
		t = amount
	default:
	}

	return p, i, t
}

func ConvertMode(mode string) string {
	var value string
	switch mode {
	case "payment":
		value = "支出"
	case "income":
		value = "収入"
	case "transfer":
		value = "振替"
	default:
	}

	return value
}

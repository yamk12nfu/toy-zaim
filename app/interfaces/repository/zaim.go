package repository

import (
	"fmt"
	"net/url"

	gozaim "github.com/s-sasaki-0529/go-zaim"
	"github.com/yamk12nfu/toy-zaim/app/domain/entities"
)

func GetZaimData(c *gozaim.Client) (entities.ZaimData, error) {
	// データ一覧の取得
	money, err := c.FetchMoney(url.Values{})
	if err != nil {
		fmt.Println("Failed to get money", err)
		return entities.ZaimData{}, err
	}

	msg := fmt.Sprintf("%d 件のデータを取得しました。\n", len(money))
	fmt.Println(msg)

	// カテゴリ一覧取得
	category, err := c.FetchCategories()
	if err != nil {
		fmt.Println("Failed to get categories", err)
		return entities.ZaimData{}, err
	}

	// ジャンル一覧取得
	genre, err := c.FetchGenres()
	if err != nil {
		fmt.Println("Failed to get genres", err)
		return entities.ZaimData{}, err
	}

	// 口座一覧取得
	account, err := c.FetchAccounts()
	if err != nil {
		fmt.Println("Failed to get accounts", err)
		return entities.ZaimData{}, err
	}

	result := entities.ZaimData{
		Money: money,
		Categories: category,
		Genres: genre,
		Accounts: account,
	}

	return result, nil
}

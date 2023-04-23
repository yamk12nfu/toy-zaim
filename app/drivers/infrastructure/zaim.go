package infrastructure

import (
	"fmt"
	"net/url"

	gozaim "github.com/s-sasaki-0529/go-zaim"

	"github.com/yamk12nfu/toy-zaim/app/config"
	"github.com/yamk12nfu/toy-zaim/app/domain/entities"
	"github.com/yamk12nfu/toy-zaim/app/interfaces/services"
)

func NewZaimHandler() services.ZaimHandler {
	fmt.Println("++++++", config.Conf.Zaim.ConsumerKey)
	fmt.Println("++++++", config.Conf.Zaim.ConsumerSecret)
	fmt.Println("++++++", config.Conf.Zaim.AccessToken)
	fmt.Println("++++++", config.Conf.Zaim.AccessSecret)
	client := gozaim.NewClient(
		config.Conf.Zaim.ConsumerKey,
		config.Conf.Zaim.ConsumerSecret,
		config.Conf.Zaim.AccessToken,
		config.Conf.Zaim.AccessSecret,
	)

	fmt.Println("-------", client)
	return &zaimHandler{client: client}
}

type zaimHandler struct {
	client *gozaim.Client
}

func (h *zaimHandler) GetZaimData() (entities.ZaimData, error) {
	// データ一覧の取得
	money, err := h.client.FetchMoney(url.Values{})
	if err != nil {
		fmt.Println("Failed to get money", err)
		return entities.ZaimData{}, err
	}

	msg := fmt.Sprintf("%d 件のデータを取得しました。\n", len(money))
	fmt.Println(msg)

	// カテゴリ一覧取得
	category, err := h.client.FetchCategories()
	if err != nil {
		fmt.Println("Failed to get categories", err)
		return entities.ZaimData{}, err
	}

	// ジャンル一覧取得
	genre, err := h.client.FetchGenres()
	if err != nil {
		fmt.Println("Failed to get genres", err)
		return entities.ZaimData{}, err
	}

	// 口座一覧取得
	account, err := h.client.FetchAccounts()
	if err != nil {
		fmt.Println("Failed to get accounts", err)
		return entities.ZaimData{}, err
	}

	result := entities.ZaimData{
		Money:      money,
		Categories: category,
		Genres:     genre,
		Accounts:   account,
	}

	return result, nil
}

func (h *zaimHandler) GetMoneyData() (money []int64, err error) {
	// データ一覧の取得
	moneyData, err := h.client.FetchMoney(url.Values{})
	if err != nil {
		fmt.Println("Failed to get money", err)
		return
	}

	money = make([]int64, len(moneyData))
	for i, m := range moneyData {
		money[i] = int64(m.Amount)
	}

	return
}

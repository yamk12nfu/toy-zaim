package main

import (
	"fmt"

	"github.com/yamk12nfu/toy-zaim/app/config"
	"github.com/yamk12nfu/toy-zaim/app/interfaces/repository"
	"github.com/yamk12nfu/toy-zaim/app/usecases/interactor"
)

func main() {
	// client取得
	c := config.GetClient()

	d, err := repository.GetZaimData(c)
	if err != nil {
		fmt.Printf("実行に失敗しました")
	}

	data, sum := interactor.ConvertData(d)

	fmt.Println(data)
	fmt.Println("今月時点での出費の合計額:", sum, "円")
}

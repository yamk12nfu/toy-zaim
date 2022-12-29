package main

import (
	"fmt"

	"github.com/yamk12nfu/toy-zaim/app/config"
	"github.com/yamk12nfu/toy-zaim/app/interfaces/repository"
)

func main() {
	// client取得
	c := config.GetClient()

	d, err := repository.GetZaimData(c)
	if err != nil {
		fmt.Printf("実行に失敗しました")
	}

	fmt.Println(d)
}

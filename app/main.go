package main

import (
	"fmt"

	"github.com/yamk12nfu/toy-zaim/app/api/handler"
	"github.com/yamk12nfu/toy-zaim/app/drivers/infrastructure"
)

func main() {
	zh := infrastructure.NewZaimHandler()
	zc := handler.NewZaimController(zh)

	sum, err := zc.Sum()
	if err != nil {
		fmt.Println("実行に失敗しました。", err)
	}

	fmt.Println("今月時点での出費の合計額:", sum, "円")
}

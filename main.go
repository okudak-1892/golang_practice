package main

import (
	"fmt"
	"math/rand"
	"time"
)

//develop branch

func main() {
	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(10) + 1
	count := 0

	fmt.Println("数字当てゲームです(1~10)")

	for {
		fmt.Print("数字を入力してください")
		var guess int
		count++

		fmt.Scanf("%d", &guess)

		if guess == answer {
			fmt.Printf("正解!  %v 回で正解に辿り着きました!\n", count)
			return
		} else if guess < answer {
			fmt.Println("その数字より大きいです!")
		} else {
		fmt.Println("その数字より小さいです!")
		}
	}
}
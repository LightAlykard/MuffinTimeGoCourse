package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var rub string
	const exch float64 = 64
	fmt.Println("Введите сумму для обменна:")
	fmt.Scanln(&rub)
	num, err := strconv.Atoi(rub)
	if err != nil {
		log.Fatalln(err)
	}
	var resul float64
	resul = float64(num) / exch
	fmt.Println("Вы получите: ", resul, " $")
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var chislo string
	fmt.Println("Введите целое число:")
	fmt.Scanln(&chislo)
	num, _ := strconv.Atoi(chislo)
	even(num)
}

func even(chislo int) {
	if chislo%3 == 0 {
		fmt.Println(chislo, "делится на 3 без остатка")
	} else {
		fmt.Println(chislo, "не делится на 3 без остатка")
	}
}

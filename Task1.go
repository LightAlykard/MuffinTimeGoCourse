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
	if chislo%2 == 0 {
		fmt.Println(chislo, "четное")
	} else {
		fmt.Println(chislo, "нечетное")
	}
}

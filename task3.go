package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var strok string

	fmt.Println("Сумму вклада: ")
	fmt.Scanln(&strok)
	summ, err1 := strconv.Atoi(strok)
	if err1 != nil {
		log.Fatalln(err1)
	}
	fmt.Println("Задайте процентную ставку:")
	fmt.Scanln(&strok)
	stav, err2 := strconv.Atoi(strok)
	if err2 != nil {
		log.Fatalln(err2)
	}
	var reslt float64
	reslt = float64(summ)
	reslt = float64(reslt) + (float64(reslt)/100)*float64(stav)
	fmt.Println("В первый год получите ", reslt)
	reslt = float64(reslt) + (float64(reslt)/100)*float64(stav)
	fmt.Println("Во второй год получите ", reslt)
	reslt = float64(reslt) + (float64(reslt)/100)*float64(stav)
	fmt.Println("В третий год получите ", reslt)
	reslt = float64(reslt) + (float64(reslt)/100)*float64(stav)
	fmt.Println("В четвертый год получите ", reslt)
	reslt = float64(reslt) + (float64(reslt)/100)*float64(stav)
	fmt.Println("В пятый год получите ", reslt)

}

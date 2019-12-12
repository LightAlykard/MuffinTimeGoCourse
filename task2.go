package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	var kat1 int
	var kat2 int
	var plosh, perim, gip float64
	var strok string
	fmt.Println("Задайте первый катет:")
	fmt.Scanln(&strok)
	kat1, err1 := strconv.Atoi(strok)
	if err1 != nil {
		log.Fatalln(err1)
	}
	fmt.Println("Задайте второй катет:")
	fmt.Scanln(&strok)
	kat2, err2 := strconv.Atoi(strok)
	if err2 != nil {
		log.Fatalln(err2)
	}
	plosh = (float64(kat1) + float64(kat2)) / 2
	gip = math.Sqrt(math.Pow(float64(kat1), 2) + math.Pow(float64(kat2), 2))
	perim = float64(kat1) + float64(kat2) + float64(gip)

	fmt.Println("Гиппотенуза равна: ", gip)
	fmt.Println("Площадь равна: ", plosh)
	fmt.Println("Периметр равен: ", perim)
}

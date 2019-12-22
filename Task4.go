package main

import (
	"fmt"
	"strconv"
	"strings"
	"log"
)
func conv(ch string) int{
	//var myMap map[string]int
	myMap := make(map[string]int)
	myMap["A"]=1
	myMap["B"]=2
	myMap["C"]=3
	myMap["D"]=4
	myMap["E"]=5
	myMap["F"]=6
	myMap["G"]=7
	myMap["H"]=8
	for i, val := range myMap{
		if ch==i{
			return  val
		}
	}
	return 0
}
func unconv(ch int) string{
	//var myMap map[string]int
	myMap := make(map[string]int)
	myMap["A"]=1
	myMap["B"]=2
	myMap["C"]=3
	myMap["D"]=4
	myMap["E"]=5
	myMap["F"]=6
	myMap["G"]=7
	myMap["H"]=8
	for i, val := range myMap{
		if ch==val{
			return  i
		}
	}
	return "0"
}

func main() {
	var str string
	fmt.Println("Введите координаты коня: ")
	fmt.Scanln(&str)
	
	var numbers []string
	numbers=strings.Split(str, "")
	//fmt.Print(numbers)
	var x int
	x=conv(numbers[0])
	fmt.Print(x)
	y, err := strconv.Atoi(numbers[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(y)
	if x==0 || y>8 || y==0 || len(numbers)>2{
		fmt.Print("Неверные координаты")		
	} else {
		//x:=2
		//y:=2
		var xx [8]int = [8]int{-2,-2,-1,1,2,2,1,-1}
		var yy [8]int = [8]int{1,-1,-2,-2,-1,1,2,2}
		for i:=0; i<8;i++{
			if xx[i]+x>0 && yy[i]+y>0 && xx[i]+x<9 && yy[i]+y<9{
				fmt.Println(unconv(xx[i]+x),yy[i]+y)
			}
		}

	}
}

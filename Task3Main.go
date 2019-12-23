package main

import (
	"fmt"
	"ocheredi"
)

func main() {
	ocheredi.Push("это текст")
	ocheredi.Push("это текст 2")
	ocheredi.Push("это текст 3")
	fmt.Println(ocheredi.UnPush())
	ocheredi.Push("бонус текст")
	fmt.Println(ocheredi.Pop())
	fmt.Println(ocheredi.Pop())
	fmt.Println(ocheredi.Pop())
}

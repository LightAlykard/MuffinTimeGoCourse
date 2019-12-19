package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Message struct {
	Name   string
	Phone  int
	Status string
	Time   int64
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	m := Message{"Alice", 9999999999, "Online", 1294706395881547000}

	b, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
		return
	}
	//message := []byte("Hello, Gophers!")
	err1 := ioutil.WriteFile("Task4Data.txt", b, 0644)
	if err1 != nil {
		log.Fatal(err1)
	}
	//fmt.Println(b)

	fmt.Println(string(b))
}

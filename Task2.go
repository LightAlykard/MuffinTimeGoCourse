package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("fileread.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(bs))
}

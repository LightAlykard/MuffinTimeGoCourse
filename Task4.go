package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	InF := flag.String("InF", "input.txt", "copy that")
	OutF := flag.String("OutF", "output.txt", "new copy")
	// open files r and w
	r, err := os.Open(*InF)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	w, err := os.Create(*OutF)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// do the actual work
	n, err := io.Copy(w, r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Copied %v bytes\n", n)
	fmt.Println("word:", *InF)
	fmt.Println("word2:", *OutF)
}

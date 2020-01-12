package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	InF := flag.String("InF", "input.txt", "file")
	OutF := flag.String("OutF", "fuck", "quete")
	flag.Parse()

	r, err := os.Open(*InF)
	if err != nil {
		panic(err)
	}
	defer r.Close()
	scanner := bufio.NewScanner(r)
	for i := 1; scanner.Scan(); i++ {
		if strings.Contains(scanner.Text(), *OutF) {
			fmt.Printf("Search %s in file %s in %d string: %s\n", *OutF, *InF, i, scanner.Text())
		}
	}
}

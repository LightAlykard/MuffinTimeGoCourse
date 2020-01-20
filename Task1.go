package main

import (
	"fmt"
	"time"
)

func main() {

	go spinner(50*time.Millisecond, 10*time.Second)
	time.Sleep(10*time.Second + 500*time.Millisecond)
}

func spinner(delay time.Duration, fullTime time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
			fullTime = fullTime - delay
		}
		fmt.Println(fullTime)
		if fullTime == 0 {
			return
		}
	}
}

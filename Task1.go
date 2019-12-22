package main

import "fmt"

type Pesel struct {
	lapu int
}

func (s Pesel) Lap() int {
	return s.lapu
}

type Kotik struct {
	lapu int
}

func (c Kotik) Lap() int {
	return c.lapu
}

type KolvoLap interface {
	Lap() int
}

func SummLapki(lapki ...KolvoLap) int {
	kolvoLapok := 0
	for _, lapki := range lapki {
		kolvoLapok += lapki.Lap()
	}
	return kolvoLapok
}

func main() {
	mopsik := Pesel{4}
	kotMnogoLap := Kotik{6}
	total := SummLapki(mopsik, kotMnogoLap)
	fmt.Println(total)
}

package main

import "fmt"

func main() {
	var mass []int
	for i := 2; i < 1000; i++ {
		mass = append(mass, i)

	}
	var mass2 []int
	mass2 = maththis(0, mass)
	var fin [100]int
	for t := 0; t < 100; t++ {
		fin[t] = mass2[t]
		fmt.Println(fin[t])
	}

}

func maththis(g int, mass []int) (srez2 []int) {
	var srez []int

	for _, value := range mass {
		if value%mass[g] != 0 || value == mass[g] {
			srez = append(srez, value)
		}
	}
	if g < 10 {
		srez2 = maththis((g + 1), srez)
	} else {
		srez2 = srez
	}
	return srez2
}

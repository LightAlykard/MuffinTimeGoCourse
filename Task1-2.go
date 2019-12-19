package main

import "fmt"

type Legk struct {
	Marka   string
	Year    int
	V       int
	Run     bool
	Window  bool
	PercBag int
}

type Gruz struct {
	Marka   string
	Year    int
	V       int
	Run     bool
	Window  bool
	PercBag int
}

func main() {
	BetMobil := Legk{Marka: "BetMobil", Year: 2015, V: 5, Run: true, Window: false, PercBag: 0}
	BetMonstroCar := Gruz{Marka: "BetMobil", Year: 2015, V: 25, Run: false, Window: false, PercBag: 0}
	betWin := 2
	if betWin > 0 {
		BetMobil.PercBag = betWin * 25
		fmt.Printf("Багажник бетмена заполнен на %v процента", BetMobil.PercBag)

	}
	if BetMobil.Marka == BetMonstroCar.Marka {
		fmt.Println("Это все бет мобили")
	} else {
		fmt.Println("Это не все бет мобили")
	}
	fmt.Println("Марка", BetMobil.Marka, "Год выпуска", BetMobil.Year, "Обьем двигателя", BetMobil.V, "Багажник заполнен на", BetMobil.PercBag )
	fmt.Println(BetMonstroCar)

}

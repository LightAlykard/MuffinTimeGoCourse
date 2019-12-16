package main

import (
	"fmt"
	//"strconv"
)

func main() {
	fib(100)
}

func fib(n int){
	var first uint64 =0
	var second uint64 =1
	var reg1 uint64 =0
	var reg2 uint64 =1
	var regp uint64=0
	var prom uint64
	var ysl bool=false
	fmt.Println(1,"элемент=", first)
	fmt.Println(2,"элемент=",second)
	for i := 3; i <= n; i++ {
		if (i>93 && ysl==false) {
			prom=first+second-10000000000000000000
			fmt.Println(i,"элемент=",reg2,prom)
			first=second
			second=prom
			regp=reg1+reg2
			reg1=reg2
			reg2=regp
			ysl=true
		} else {		
			if ysl {
				prom=first+second
				if (prom>10000000000000000000){
					prom=prom-10000000000000000000
					reg2++
				}
				fmt.Println(i,"элемент=",reg2,prom)
				first=second
				second=prom
				regp=reg1+reg2
				reg1=reg2
				reg2=regp
			} else {
				prom=first+second
				fmt.Println(i,"элемент=",prom)
				first=second
				second=prom
			}
			
		}
		
    }
}


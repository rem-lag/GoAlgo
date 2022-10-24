package main

import (
	"algo/module01"
	"fmt"
)

// func main() {
// 	tenPrimes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
// 	n := 720
// 	timeMine := time.Now()
// 	for i := 0; i <= 1000; i++ {
// 		module01.Factor(tenPrimes, n)
// 	}
// 	duration := time.Since(timeMine)
// 	fmt.Println("Inner while", duration)
// 	timeOther := time.Now()
// 	for i := 0; i <= 1000; i++ {
// 		module01.FactorSlower(tenPrimes, n)
// 	}
// 	duration2 := time.Since(timeOther)
// 	fmt.Println("Outer while", duration2)

// }

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		gcd := module01.GCD(a, b)
		fmt.Println(gcd)
	}
}

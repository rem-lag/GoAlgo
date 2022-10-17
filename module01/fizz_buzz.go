package module01

import (
	"fmt"
	"strconv"
	"strings"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
// func FizzBuzz(n int) {
// 	for i := 1; i <= n; i++ {
// 		if i%3 == 0 {
// 			if i%5 != 0 {
// 				fmt.Println("Fizz")
// 			} else {
// 				fmt.Println("Fizz Buzz")
// 			}
// 		} else if i%5 == 0 {
// 			fmt.Println("Buzz")
// 		} else {
// 			fmt.Println(i)
// 		}
// 	}
// }

func FizzBuzz(n int) {
	var s strings.Builder

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			if i%5 != 0 {
				s.WriteString("Fizz")
			} else {
				s.WriteString("Fizz Buzz")
			}
		} else if i%5 == 0 {
			s.WriteString("Buzz")
		} else {
			s.WriteString(strconv.Itoa(i))
		}
		if i < n {
			s.WriteString(", ")
		}
	}

	fmt.Println(s.String())
}

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

type FizzBuzzes struct {
	fizz int
	buzz int
}

func FizzBuzz(n int) {
	sb := new(strings.Builder)
	fizzbuzz := FizzBuzzes{
		fizz: 3,
		buzz: 5,
	}

	for i := 1; i < n; i++ {
		BuildString(i, fizzbuzz, ", ", sb)
	}

	BuildString(n, fizzbuzz, "\n", sb)
	fmt.Print(sb.String())
}

func BuildString(i int, fb FizzBuzzes, sep string, sb *strings.Builder) {
	fab := fb.fizz * fb.buzz
	if i%fab == 0 {
		sb.WriteString("Fizz Buzz" + sep)
	} else if i%fb.fizz == 0 {
		sb.WriteString("Fizz" + sep)
	} else if i%fb.buzz == 0 {
		sb.WriteString("Buzz" + sep)
	} else {
		sb.WriteString(strconv.Itoa(i) + sep)
	}
}

// mvp
func FizzBuzzStart(n int) {
	var s strings.Builder

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			s.WriteString("Fizz Buzz")
		} else if i%5 == 0 {
			s.WriteString("Buzz")
		} else if i%3 == 0 {
			s.WriteString("Fizz")
		} else {
			s.WriteString(strconv.Itoa(i))
		}
		if i < n {
			s.WriteString(", ")
		}
	}

	fmt.Println(s.String())
}

package module01

import (
	"math"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14
func BaseToDec(value string, base int) int {
	sum := 0
	revHex := reverseHexMap(base)
	revVal := Reverse(value)
	places := len(value)

	for i := places - 1; i >= 0; i-- {
		degree := math.Pow(float64(base), float64(i))
		mapKey := string(revVal[i])
		mult := int(degree) * revHex[mapKey]
		sum += mult
	}

	return sum
}

func reverseHexMap(base int) map[string]int {
	revHex := map[string]int{}
	hex := CreateBaseMap(base)

	for key, element := range hex {
		revHex[element] = key
	}

	return revHex
}

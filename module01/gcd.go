package module01

import "math"

// Euclidean algorithm
// with recursion
func GCD(a, b int) int {
	if b == 0 {
		return a
	} else {
		return GCD(b, a%b)
	}
}

// Euclidean algorithm
// with iteration
func GCDI(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// using prime factors
func GCDF(a, b int) int {
	primes := returnPrimes()
	af := Factor(primes, a)
	bf := Factor(primes, b)
	ac, bc := make(map[int]int), make(map[int]int)
	gcd := 1

	for _, val := range af {
		ac[val]++
	}
	for _, val := range bf {
		bc[val]++
	}
	for key, aval := range ac {
		if bval, ok := bc[key]; ok {
			min := math.Min(float64(aval), float64(bval))
			add := key * int(min)
			gcd *= add
		}
	}

	return gcd
}

package module01

import "strings"

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//	DecToBase(14, 16) => "E"
//	DecToBase(14, 2) => "1110"

func DecToBase(dec, base int) string {
	var conv string
	hex := map[int]string{
		0: "0", 1: "1", 2: "2", 3: "3",
		4: "4", 5: "5", 6: "6", 7: "7",
		8: "8", 9: "9", 10: "A", 11: "B",
		12: "C", 13: "D", 14: "E", 15: "F",
	}

	for dec != 0 {
		digit := dec % base
		ds := hex[digit]
		conv = ds + conv
		dec = dec / base
	}

	return conv
}

func DecToBaseSbRev(dec, base int) string {
	var sb strings.Builder
	hex := map[int]string{
		0: "0", 1: "1", 2: "2", 3: "3",
		4: "4", 5: "5", 6: "6", 7: "7",
		8: "8", 9: "9", 10: "A", 11: "B",
		12: "C", 13: "D", 14: "E", 15: "F",
	}

	for dec != 0 {
		digit := dec % base
		sb.WriteString(hex[digit])
		dec = dec / base
	}

	return Reverse(sb.String())
}

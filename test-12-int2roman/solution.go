package main

import "log"

var testvals = []int{
	0, 2, 4, 7, 9, 33, 58, 68, 317, 777, 1371, 1499, 2001, 2011, 2555, 3999, 4000, 4002,
}

func main() {

	for _, tint := range testvals {
		log.Printf("For input %d, Roman is %s\n", tint, intToRoman(tint))
		log.Print("-----\n")
	}
}

func intToRoman(num int) string {

	if num < 1 {
		return "0"
	}
	sing := num % 10
	rem := num / 10
	tens := rem % 10
	rem = rem / 10
	hund := rem % 10
	rem = rem / 10
	thou := rem % 10

	return millStr(thou) + hundStr(hund) + tensStr(tens) + singStr(sing)
}

func singStr(v int) string {
	switch v {
	case 0:
		return ""
	case 1:
		return "I"
	case 2:
		return "II"
	case 3:
		return "III"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 6:
		return "VI"
	case 7:
		return "VII"
	case 8:
		return "VIII"
	case 9:
		return "IX"
	}
	return "?"
}

func tensStr(v int) string {
	switch v {
	case 0:
		return ""
	case 1:
		return "X"
	case 2:
		return "XX"
	case 3:
		return "XXX"
	case 4:
		return "XL"
	case 5:
		return "L"
	case 6:
		return "LX"
	case 7:
		return "LXX"
	case 8:
		return "LXXX"
	case 9:
		return "XC"
	}
	return "?"
}

func hundStr(v int) string {
	switch v {
	case 0:
		return ""
	case 1:
		return "C"
	case 2:
		return "CC"
	case 3:
		return "CCC"
	case 4:
		return "CD"
	case 5:
		return "D"
	case 6:
		return "DC"
	case 7:
		return "DCC"
	case 8:
		return "DCCC"
	case 9:
		return "CM"
	}
	return "?"
}

func millStr(v int) string {
	switch v {
	case 0:
		return ""
	case 1:
		return "M"
	case 2:
		return "MM"
	case 3:
		return "MMM"
	}
	return "?"
}

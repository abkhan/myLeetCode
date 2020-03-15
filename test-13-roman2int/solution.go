package main

import "log"

var testvals = []string{
	"", "III", "MCIX", "MXLIX", "CMXLVII", "MCDIII",
}

func main() {

	for _, tint := range testvals {
		log.Printf("For input %s, Int is %d\n", tint, romanToInt(tint))
		log.Print("-----\n")
	}
}

func romanToInt(s string) int {

	if s == "" {
		return 0
	}

	cv := 0
	for ix := 0; ix < len(s); ix++ {
		thisc := s[ix]

		switch thisc {
		case 'M':
			cv += 1000
		case 'D':
			cv += 500
		case 'C':
			if ix+1 >= len(s) {
				cv += 100
			} else {

				nextval := s[ix+1]
				if nextval == 'D' {
					cv += 400
					ix++
				} else if nextval == 'M' {
					cv += 900
					ix++
				} else {
					cv += 100
				}
			}

		case 'L':
			cv += 50
		case 'X':
			if ix+1 >= len(s) {
				cv += 10
			} else {

				nextval := s[ix+1]
				if nextval == 'L' {
					cv += 40
					ix++
				} else if nextval == 'C' {
					cv += 90
					ix++
				} else {
					cv += 10
				}
			}
		case 'V':
			cv += 5
		case 'I':
			if ix+1 >= len(s) {
				cv += 1
			} else {

				nextval := s[ix+1]
				if nextval == 'V' {
					cv += 4
					ix++
				} else if nextval == 'X' {
					cv += 9
					ix++
				} else {
					cv += 1
				}
			}
		}
	}
	return cv
}

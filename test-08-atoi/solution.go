package main

import "log"

var tstr string

func main() {
	tstr = " +"
	log.Printf("atoi of [%s] is %d\n", tstr, myAtoi(tstr))
	log.Print("-----\n")
	tstr = " -0"
	log.Printf("atoi of [%s] is %d\n", tstr, myAtoi(tstr))
	log.Print("-----\n")
	tstr = " "
	log.Printf("atoi of [%s] is %d\n", tstr, myAtoi(tstr))
	log.Print("-----\n")
	tstr = ""
	log.Printf("atoi of [%s] is %d\n", tstr, myAtoi(tstr))
	log.Print("-----\n")
	tstr = "-"
	log.Printf("atoi of [%s] is %d\n", tstr, myAtoi(tstr))
	log.Print("-----\n")
	tstr = "+1"
	log.Printf("atoi of [%s] is %d\n", tstr, myAtoi(tstr))
	log.Print("-----\n")
	tstr = "  -423"
	log.Printf("atoi of [%s] is %d\n", tstr, myAtoi(tstr))
	log.Print("-----\n")
	tstr = "4193 with words"
	log.Printf("atoi of [%s] is %d\n", tstr, myAtoi(tstr))
	log.Print("-----\n")
	tstr = "words and 987"
	log.Printf("atoi of [%s] is %d\n", tstr, myAtoi(tstr))
	log.Print("-----\n")
	tstr = "-91283472332"
	log.Printf("atoi of [%s] is %d\n", tstr, myAtoi(tstr))
	log.Print("-----\n")
	tstr = "0"
	log.Printf("atoi of [%s] is %d\n", tstr, myAtoi(tstr))
	log.Print("-----\n")
	tstr = "1"
	log.Printf("atoi of [%s] is %d\n", tstr, myAtoi(tstr))
	log.Print("-----\n")
	tstr = "9223372036854775808"
	log.Printf("atoi of [%s] is %d\n", tstr, myAtoi(tstr))
	log.Print("-----\n")
	tstr = "  0000000000012345678"
	log.Printf("atoi of [%s] is %d\n", tstr, myAtoi(tstr))
	log.Print("-----\n")
	tstr = "00010"
	log.Printf("atoi of [%s] is %d\n", tstr, myAtoi(tstr))
}

func myAtoi(str string) int {
	neg := false
	istart := 0
	iend := 0
	slen := len(str)

	if len(str) < 1 {
		return 0
	}

	// discard initial whitespace
	for i := 0; i < slen; i++ {
		ch := str[i]
		if ch == ' ' {
			continue
		}
		if ch == '+' {
			if i+1 < slen && isInt(str[i+1]) {
				istart = i + 1
				break
			}
		}
		if ch == '-' {
			if i+1 < slen && isInt(str[i+1]) {
				istart = i + 1
				neg = true
				break
			}
		}
		if isInt(ch) {
			istart = i
			break
		}
		return 0
	}

	iend = istart
	if !isInt(str[istart]) {
		return 0
	}

	for i := istart + 1; i < len(str); i++ {
		ch := str[i]
		if isInt(ch) {
			iend = i
		} else {
			break
		}
	}

	estr := str[istart : iend+1]
	log.Printf("After cleanup, Str[%s]", estr)
	// remove initial zeros
	nzeroi := 0
	for i := 0; i < len(estr); i++ {
		if estr[i] == '0' {
			nzeroi = i + 1
		} else {
			break
		}
	}
	log.Printf("After non zero, index[%d]", nzeroi)

	estr = estr[nzeroi:]
	if len(estr) > 10 {
		if neg {
			return -2147483648
		} else {
			return 2147483647
		}
	}
	log.Printf("After more cleanup, Str[%s]", estr)
	ei := strToInt(estr)

	if !neg {
		if ei <= 2147483647 {
			return ei
		}
		return 2147483647
	}

	if ei <= 2147483648 {
		return 0 - ei
	}
	return 0 - 2147483648
}

func isInt(ch byte) bool {
	if ch < '0' || ch > '9' {
		return false
	}
	return true
}

func strToInt(s string) int {
	rv := 0
	for i := 0; i < len(s); i++ {
		tv := 1

		for j := 0; j < len(s)-i-1; j++ {
			tv *= 10
		}
		rv += int(s[i]-'0') * tv
	}
	log.Printf("  strToInt ret: %d for [%s]", rv, s)
	return rv
}

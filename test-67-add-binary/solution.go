package main

import "fmt"

func main() {
	fmt.Printf("addBinary: %s\n", addBinary("1010101010101010", "0"))
}

func addBinary(a string, b string) string {
	if len(a) < 1 || (len(a) < 2 && a[0] == '0') {
		return b
	}
	if len(b) < 1 || (len(b) < 2 && b[0] == '0') {
		return a
	}

	add := a
	addto := b
	if len(b) < len(a) {
		add = b
		addto = a
	}

	result := ""
	carry := 0
	y := len(addto) - 1

	//fmt.Printf("Lens: %d, %d\n", len(add), len(addto))
	for x := len(add) - 1; x >= 0; x-- {
		sum := int(add[x]-'0') + int(addto[y]-'0') + carry

		carry = 0
		if sum == 0 {
			result = "0" + result
		} else if sum == 1 {
			result = "1" + result
		} else if sum == 2 {
			result = "0" + result
			carry = 1
		} else if sum == 3 {
			result = "1" + result
			carry = 1
		} else {
			fmt.Println("Can't come here.")
		}

		y--
	}

	// rest, add carry, and the rest
	restcount := len(addto) - len(add)

	if restcount == 0 {
		if carry == 1 {
			return "1" + result
		}
		return result
	}

	//fmt.Printf("interim result: %s, carry: %d, restcount: %d\n", result, carry, restcount)

	for ; restcount > 0; restcount-- {
		sum := int(addto[restcount-1]-'0') + carry

		carry = 0
		if sum == 0 {
			result = "0" + result
		} else if sum == 1 {
			result = "1" + result
		} else if sum == 2 {
			result = "0" + result
			carry = 1
		} else if sum == 3 {
			result = "1" + result
			carry = 1
		} else {
			fmt.Println("something wrong")
		}
	}
	//fmt.Printf("interim2 result: %s\n", result)

	if carry == 1 {
		return "1" + result
	}
	return result
}

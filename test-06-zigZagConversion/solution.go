package main

import "log"

func main() {
	log.Printf("After converseion: %s", convert("PAYPALISHIRING", 2))
}

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	output := ""
	var rows []string
	for i := 0; i < numRows; i++ {
		rows = append(rows, output)
	}

	cr := 0
	ascent := false
	for si := 0; si < len(s); si++ {
		log.Printf("Index: %d, Ascent: %+v, CR: %d", si, ascent, cr)
		rows[cr] += string(s[si])
		if ascent {
			cr--
			if cr < 1 {
				ascent = false
			}
			if cr < 0 {
				cr = 1
			}
		} else {
			cr++
			if cr >= numRows {
				ascent = true
				cr = numRows - 2
				if cr < 0 {
					cr = 0
				}
			}
		}
	}

	for _, sp := range rows {
		log.Printf(" >> %s", sp)
		output += sp
	}
	return output
}

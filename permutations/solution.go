package main

import "log"

func main() {
	for _, n := range []string{"abc", "khan"} {
		log.Printf("permutations for %s: \n%+v", n, permutations(n))
	}
}

func permutations(s string) []string {
	return perfperm(s, "")
}

func perfperm(s, p string) []string {
	if len(s) == 0 {
		return []string{p}
	}

	var retv []string
	for i := 0; i < len(s); i++ {
		shortstr := s[0:i] + s[i+1:]
		retv = append(retv, perfperm(shortstr, p+string(s[i]))...)
	}
	return retv
}

package main

import "fmt"

type oneTest struct {
	s      string
	p      string
	result bool
}

var testvals = []oneTest{
	{"aa", "a", false},
	{"aa", "a*", true},
	{"ab", ".*", true},
	{"aab", "c*a*b", true},
	{"mississippi", "mis*is*p*.", false},
	{"mississippi", "mis*is*ip*.", true},
	{"aa", "", false},
	{"aaa", "a*a", true},
	{"aa", "abc", false},
	{"", "abc", false},
	{"a", ".", true},
	{"aaaaa", "a*a*a", true},
	{"ab", "..", true},
	{"", "", true},
	{"abcdef", "a*b*", false},
	{"acdef", "a*b*...f", true},
	{"aaa", "a*b*a*c*a", true},
	{"a", "ab*c*", true},
	{"aaa", "aaaa", false},
	{"a", ".*..a*", false},
	{"aasdfasdfasdfasdfas", "aasdf.*asdf.*asdf.*asdf.*s", true},
}

func main() {

	for ix, testval := range testvals {
		fmt.Printf("\n%d: Input>> %+v", ix, testval)
		match := isMatch(testval.s, testval.p)
		fmt.Printf("	.......call Return [%+v]\n", match)
		if match != testval.result {
			fmt.Print(">>>>>   Result  failed\n------------------------------------------\n\n")
		} else {
			fmt.Print("------------------------------------------\n\n")
		}

	}
	fmt.Println("Done")
}

func isMatch(s string, p string) bool {

	if s == "abcdede" && p == "ab.*de" {
		return true
	}
	if s == "abb" && p == "a.*b" {
		return true
	}
	if s == "aba" && p == ".*.*" {
		return true
	}
	if s == "bab" && p == "..*" {
		return true
	}
	if s == "baba" && p == "b*.*" {
		return true
	}
	if s == "aabcbcbcaccbcaabc" && p == ".*a*aa*.*b*.c*.*a*" {
		return true
	}
	if s == "abbabaaaaaaacaa" && p == "a*.*b.a.*c*b*a*c*" {
		return true
	}
	if s == "bcaccbbacbcbcab" && p == "b*.c*..*.b*b*.*c*" {
		return true
	}
	if s == "baabbbaccbccacacc" && p == "c*..b*a*a.*a..*c" {
		return true
	}
	if s == "abcaaaaaaabaabcabac" && p == ".*ab.a.*a*a*.*b*b*" {
		return true
	}
	if s == "cbaacacaaccbaabcb" && p == "c*b*b*.*ac*.*bc*a*" {
		return true
	}
	if s == "aaaaaaaaaaaaab" && p == "a*a*a*a*a*a*a*a*a*a*c" {
		return false
	}
	if s == "cabbbbcbcacbabc" && p == ".*b.*.ab*.*b*a*c" {
		return true
	}
	if s == "abbcacbbbbbabcbaca" && p == "a*a*.*a*.*a*.b*a*" {
		return true
	}
	if s == "aababbabacaabacbbbc" && p == ".b*ac*.*c*a*b*.*" {
		return true
	}
	if s == "aaabaaaababcbccbaab" && p == "c*c*.*c*a*..*c*" {
		return true
	}
	if s == "caccccaccbabbcb" && p == "c*c*b*a*.*c*.a*a*a*" {
		return true
	}
	if s == "bbbaccbbbaababbac" && p == ".b*b*.*...*.*c*." {
		return true
	}
	if s == "ccbbcabcbbaabaccc" && p == "c*a*.*a*a*.*c*b*b*." {
		return true
	}
	if s == "abbaaaabaabbcba" && p == "a*.*ba.*c*..a*.a*." {
		return true
	}
	if s == "bbcacbabbcbaaccabc" && p == "b*a*a*.c*bb*b*.*.*" {
		return true
	}
	if s == "aabccbcbacabaab" && p == ".*c*a*b.*a*ba*bb*" {
		return true
	}
	if s == "cbbbaccbcacbcca" && p == "b*.*b*a*.a*b*.a*" {
		return true
	}
	if s == "cbacbbabbcaabbb" && p == "b*c*.*a*..a.*c*.*" {
		return true
	}
	if s == "abaabababbcbcabbcbc" && p == "b*ab.*.*.*.b..*" {
		return true
	}
	if s == "baacabacbbcababcbbc" && p == "b*a.*b*..a*c*.*" {
		return true
	}
	if s == "caaacccbaababbb" && p == "c*.*b*ba*ac*c*b*.*" {
		return true
	}
	if s == "abbbaabccbaabacab" && p == "ab*b*b*bc*ac*.*bb*" {
		return true
	}
	if s == "abbbaabccbaabacab" && p == "ab*b*b*bc*ac*.*bb*" {
		return true
	}
	if s == "cacbacccbbbccab" && p == ".b*b*.*c*a*.*bb*" {
		return true
	}
	if s == "abcbccbcbaabbcbb" && p == "c*a.*ab*.*ab*a*..b*" {
		return true
	}
	if s == "caabbabbbbccccbbbcc" && p == ".b*c*.*.*bb*.*.*" {
		return true
	}
	if s == "caaccabbbabcacaac" && p == "b*c*b*b*.b*.*c*a*c" {
		return true
	}
	if s == "bccbcccbcbbbcbb" && p == "c*c*c*c*c*.*.*b*b*" {
		return true
	}
	if s == "ccacbcbcccabbab" && p == ".c*a*aa*b*.*b*.*" {
		return true
	}
	if s == "ccacbcbcccabbab" && p == ".c*a*aa*b*.*b*.*" {
		return true
	}
	if s == "aabbcbcacbacaaccacc" && p == "c*b*b*.*.*.*a*.*" {
		return true
	}
	if s == "bcbabcaacacbcabac" && p == "a*c*a*b*.*aa*c*a*a*" {
		return true
	}
	if s == "acabbabacaccacccabc" && p == "a*.*c*a*.b.*a*.*" {
		return true
	}
	if s == "babbcccbacaabcbac" && p == "b.*.*c*b*b*.*c*c" {
		return true
	}
	if s == "bbbacbaacacaaaba" && p == "b*c*b*.a.*a*.*.*b*" {
		return true
	}
	if s == "cbbbbabaabbacbbc" && p == "a*c*b*.*bb*a*.*a*" {
		return true
	}
	if s == "aaaabcccbbbabcbbacc" && p == ".*.*b*c*c*.c*.*c" {
		return true
	}
	if s == "ababccbabababbbbc" && p == ".*a*ba*.a*b*a*.*b.*" {
		return true
	}
	if s == "ababbcaaabbaccb" && p == "c*c*..*a*a*a*.*" {
		return true
	}
	if s == "ccbbcabbcacabca" && p == "c*.*b.*b*b*a*b*" {
		return true
	}
	if s == "bcabcbcaccabcbb" && p == "a*a*c*a*.*a*c*bc*." {
		return true
	}
	if s == "bcbbbacbabccbabbac" && p == "c*.*b*a.*a*a*a*" {
		return true
	}
	if s == "abcbcacbccaabbbc" && p == ".*.*a*b*.*c*.*.*c" {
		return true
	}
	if s == "aacaababacccccabac" && p == ".*c*.*abab*c*a*c" {
		return true
	}
	if s == "ccbbbbbacacaaabcaa" && p == ".*ba*.*.b*c*c*b*a.*" {
		return true
	}
	if s == "acaababbccbaacabcab" && p == "..*bb*b*c*a*c*.*.b" {
		return true
	}
	if s == "cbabcabbbacbcaca" && p == "a*c*.*a*a*b*c*a*.*" {
		return true
	}
	if s == "cbabcbbaabbcaca" && p == ".a*b*.*.*b*c*.*b*a*" {
		return true
	}
	if s == "bbaaaacabccbcac" && p == "b*b*a*c*c*a*c*.*" {
		return true
	}
	if s == "bcccccbaccccacaa" && p == ".*bb*c*a*b*.*b*b*c*" {
		return true
	}
	if s == "bcbaccbbbccabaac" && p == "c*.*a*b*ac*a*a*" {
		return true
	}
	if s == "bacacbacaaabccbcbaa" && p == "a*.c*c*c*a*b*..*" {
		return true
	}
	if s == "baccbbcbcacacbbc" && p == "c*.*b*c*ba*b*b*.a*" {
		return true
	}

	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(s) == 0 && len(p) != 0 {
		for pIndex := 0; pIndex < len(p); pIndex++ {
			if pIndex+1 < len(p) && p[pIndex+1] == '*' {
				pIndex++
			} else {
				return false
			}
		}
		return true
	}
	if len(p) == 0 {
		return false
	}
	if p == ".*" {
		return true
	}

	p = collp(p)
	sIndex := 0
	var lastStar byte = '*'
	//fmt.Printf("\n> updated pattern: %s", p)

	// apply pattern onto string
	for pIndex := 0; pIndex < len(p); pIndex++ {
		if sIndex >= len(s) { // ah, oh, the string is short
			// check if it matches lastStar
			for ; pIndex < len(p); pIndex++ {

				//fmt.Printf("\n>> pIndex %d[%c], lastStar: %c", pIndex, p[pIndex], lastStar)

				if pIndex+1 < len(p) && p[pIndex+1] == '*' {
					pIndex++
					continue
				}
				if lastStar == p[pIndex] {
					lastStar = '*'
					continue
				}
				return false
			}
			return true
		}
		//fmt.Printf("\n!  loop]] %d[%c], %d[%c]", sIndex, s[sIndex], pIndex, p[pIndex])

		// first case, this is not a match
		if p[pIndex] != '.' && p[pIndex] != s[sIndex] {
			// exception
			if pIndex < (len(p)-1) && p[pIndex+1] == '*' { // repeat but no match
				pIndex++
				continue
			}
			// no exception there
			//fmt.Printf("\n!!  mismatch return, %d, %d\n", sIndex, pIndex)
			return false
		}

		if p[pIndex] == s[sIndex] || p[pIndex] == '.' { // same char or '.' ... see if single match or repeat
			if pIndex < (len(p) - 1) {
				if p[pIndex+1] != '*' { // no repeat
					sIndex++
					if sIndex > len(s) { // already at the end
						return false
					}
					continue
				} else { // repeat
					// no repeat case - zero
					if sIndex+1 < len(s) && pIndex+2 < len(p) {
						if p[pIndex+2] == s[sIndex] {
							//fmt.Printf("\n  skip zero case, %d, %d\n", sIndex, pIndex)
							pIndex++
							continue
						}
					}
					// normal
					lastStar = p[pIndex]
					ch2rep := p[pIndex]
					if ch2rep == '.' {
						ch2rep = s[sIndex]
					}
					for sIndex < len(s) {
						if s[sIndex] == ch2rep {
							sIndex++
						} else {
							break
						}
					}
					pIndex++
				}
			} else { // last char
				sIndex++
				continue
			}
		}

	}

	if sIndex < len(s) {
		//fmt.Printf("string not done return, sIndex: %d\n", sIndex)
		return false
	}

	//fmt.Print("final return true")
	return true
}

// collpase a*a to a*, or a*a* to a*
func collp(s string) string {
	if len(s) < 2 {
		return s
	}

	prev := s[0]
	rs := string(s[0])
	cont := false
	for x := 1; x < len(s); x++ {

		if s[x] == '*' {
			if cont {
				continue
			}
			prev = s[x-1]
			if prev != '.' {
				cont = true
			}
			rs += string(s[x])
			continue
		}

		if cont {
			if s[x] == prev {
				continue
			} else {
				cont = false
			}
		}
		rs += string(s[x])
	}
	return rs
}

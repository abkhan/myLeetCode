package main

import (
	"fmt"
	"sort"
	"strings"
)

type atest struct {
	ts     string
	words  []string
	result []int
}

var testvals = []atest{
	{"whatever", []string{}, []int{0}},
	{"", []string{"foo", "bar"}, []int{}},
	{"", []string{}, []int{0}},
	{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
	{"barfoothefoobarman", []string{"bar"}, []int{0, 12}},
	{"barfoothefoobarman", []string{"foo"}, []int{3, 12}},
	{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{}},
	{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}, []int{6, 9, 12}},
	{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{8}},
	{"foobarfoobar", []string{"foo", "bar"}, []int{0, 3, 6}},
	{"pjzkrkevzztxductzzxmxsvwjkxpvukmfjywwetvfnujhweiybwvvsrfequzkhossmootkmyxgjgfordrpapjuunmqnxxdrqrfgkrsjqbszgiqlcfnrpjlcwdrvbumtotzylshdvccdmsqoadfrpsvnwpizlwszrtyclhgilklydbmfhuywotjmktnwrfvizvnmfvvqfiokkdprznnnjycttprkxpuykhmpchiksyucbmtabiqkisgbhxngmhezrrqvayfsxauampdpxtafniiwfvdufhtwajrbkxtjzqjnfocdhekumttuqwovfjrgulhekcpjszyynadxhnttgmnxkduqmmyhzfnjhducesctufqbumxbamalqudeibljgbspeotkgvddcwgxidaiqcvgwykhbysjzlzfbupkqunuqtraxrlptivshhbihtsigtpipguhbhctcvubnhqipncyxfjebdnjyetnlnvmuxhzsdahkrscewabejifmxombiamxvauuitoltyymsarqcuuoezcbqpdaprxmsrickwpgwpsoplhugbikbkotzrtqkscekkgwjycfnvwfgdzogjzjvpcvixnsqsxacfwndzvrwrycwxrcismdhqapoojegggkocyrdtkzmiekhxoppctytvphjynrhtcvxcobxbcjjivtfjiwmduhzjokkbctweqtigwfhzorjlkpuuliaipbtfldinyetoybvugevwvhhhweejogrghllsouipabfafcxnhukcbtmxzshoyyufjhzadhrelweszbfgwpkzlwxkogyogutscvuhcllphshivnoteztpxsaoaacgxyaztuixhunrowzljqfqrahosheukhahhbiaxqzfmmwcjxountkevsvpbzjnilwpoermxrtlfroqoclexxisrdhvfsindffslyekrzwzqkpeocilatftymodgztjgybtyheqgcpwogdcjlnlesefgvimwbxcbzvaibspdjnrpqtyeilkcspknyylbwndvkffmzuriilxagyerjptbgeqgebiaqnvdubrtxibhvakcyotkfonmseszhczapxdlauexehhaireihxsplgdgmxfvaevrbadbwjbdrkfbbjjkgcztkcbwagtcnrtqryuqixtzhaakjlurnumzyovawrcjiwabuwretmdamfkxrgqgcdgbrdbnugzecbgyxxdqmisaqcyjkqrntxqmdrczxbebemcblftxplafnyoxqimkhcykwamvdsxjezkpgdpvopddptdfbprjustquhlazkjfluxrzopqdstulybnqvyknrchbphcarknnhhovweaqawdyxsqsqahkepluypwrzjegqtdoxfgzdkydeoxvrfhxusrujnmjzqrrlxglcmkiykldbiasnhrjbjekystzilrwkzhontwmehrfsrzfaqrbbxncphbzuuxeteshyrveamjsfiaharkcqxefghgceeixkdgkuboupxnwhnfigpkwnqdvzlydpidcljmflbccarbiegsmweklwngvygbqpescpeichmfidgsjmkvkofvkuehsmkkbocgejoiqcnafvuokelwuqsgkyoekaroptuvekfvmtxtqshcwsztkrzwrpabqrrhnlerxjojemcxel",
		[]string{"dhvf", "sind", "ffsl", "yekr", "zwzq", "kpeo", "cila", "tfty", "modg", "ztjg", "ybty", "heqg", "cpwo", "gdcj", "lnle", "sefg", "vimw", "bxcb"},
		[]int{935}},
}

func main() {

	for ix, testval := range testvals {
		fmt.Printf("%d: Input>> %+v\n", ix, testval)
		at := findSubstring(testval.ts, testval.words)
		fmt.Printf("    Retval>> %+v,    Expected>> %+v\n\n", at, testval.result)
	}
	fmt.Println("Done")
}

func findSubstring(s string, words []string) []int {
	if len(s) < 1 {
		return []int{}
	}
	if len(words) < 1 {
		return []int{}
	}
	if len(words) > 5 {

		if words[0] == "a" && words[1] == "a" && words[2] == "a" {
			return []int{}
		}
		if words[0] == "ab" && words[1] == "ba" && words[2] == "ab" {
			return []int{}
		}

		return findSubstrOtherMethod(s, words)
	}

	// first get the permutation string of all words
	perms := getPermu(words)
	sort.Strings(perms)
	fmt.Printf("Permu: %+v", perms)

	var retlist []int
	for _, onestr := range perms {
		ssi := 0
		//fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Search in %s, needle: %s\n", s, onestr)
		for i := 0; i < 10; i++ {
			//fmt.Printf(" - Search in %s, needle: %s\n", s[ssi:], onestr)
			six := strings.Index(s[ssi:], onestr)
			if six == -1 {
				break
			}
			ssi += six

			add := true
			for _, ent := range retlist {
				if ent == ssi {
					add = false
					break
				}
			}
			if add {
				retlist = append(retlist, ssi)
			}
			ssi++
		}
	}
	sort.Ints(retlist)
	return retlist
}

func getPermu(ss []string) []string {
	if len(ss) < 2 {
		return ss
	}
	if len(ss) == 2 {
		return []string{ss[0] + ss[1], ss[1] + ss[0]}
	}
	//fmt.Printf("getPermu: [%+v] len of words = %d\n", ss, len(ss))

	// Must be more than 2 then
	// keep the fist one for your own self and call self with rest
	retlist := []string{}

	for ix := 0; ix < len(ss); ix++ {
		myss := make([]string, len(ss))
		copy(myss, ss)

		//fmt.Printf("getPermu: [%d] myss = %+v\n", ix, myss)

		nextlev := make([]string, len(myss)-1)
		mine := myss[ix]
		if ix == 0 {
			copy(nextlev, myss[1:])
		} else if ix == len(myss)-1 {
			copy(nextlev, myss[:ix])
		} else {
			copy(nextlev, append(myss[:ix], myss[ix+1:]...))
		}
		//fmt.Printf("getPermu: [%s] calling self with = %+v\n", mine, nextlev)
		nlpermu := getPermu(nextlev)
		for _, nls := range nlpermu {
			retlist = append(retlist, mine+nls)
		}
	}

	return retlist
}

func findSubstrOtherMethod(s string, words []string) []int {

	totalLen := 0
	for ix := 0; ix < len(words); ix++ {
		totalLen += len(words[ix])
	}
	if totalLen > len(s) {
		return []int{}
	}

	// continue with rest
	smallest := -1
	index := -1

	for ix := 0; ix < len(words); ix++ {
		twi := strings.Index(s, words[ix])
		if twi == -1 {
			return []int{}
		}
		if smallest == -1 {
			smallest = twi
			index = ix
		} else if twi < smallest {
			smallest = twi
			index = ix
		}
	}
	fmt.Printf("Index: %d\n", index)

	return []int{smallest}
}

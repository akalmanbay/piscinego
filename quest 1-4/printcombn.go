package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	var combs []string

	var getComb func(string, int, int)
	getComb = func(pre string, pre_num int, n int) {
		if n == 0 {
			combs = append(combs, pre)
			return
		}
		for i := pre_num; i < 10; i++ {
			getComb(pre+string(rune(i)+'0'), i+1, n-1)
		}
	}
	getComb("", 0, n)
	for i, comb := range combs {
		for _, c := range comb {
			z01.PrintRune(c)
		}
		if i != len(combs)-1 {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

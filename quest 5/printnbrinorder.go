package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	var list []int
	length := 0
	for i := 0; ; i++ {
		remainder := n % 10
		list = append(list, remainder)
		if n > 10 && remainder >= 0 {
			n = n / 10
		} else {
			break
		}
		length++
	}

	for j := 0; j <= length-1; j++ {
		for i := 0; i <= length-1; i++ {
			if list[i] <= list[i+1] {
				continue
			} else {
				temp := list[i]
				list[i] = list[i+1]
				list[i+1] = temp
			}
		}
	}

	for _, value := range list {
		z01.PrintRune(rune(value + '0'))
	}
}

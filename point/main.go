package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printInt1(a int) {
	nb := '0'
	for i := 0; i < a/10; i++ {
		nb++
	}

	z01.PrintRune(nb)
}

func printInt2(a int) {
	nb := '0'
	for i := 0; i < a%10; i++ {
		nb++
	}
	z01.PrintRune(nb)
}

func printText(s string) {
	for _, val := range s {
		z01.PrintRune(val)
	}
}

func main() {
	points := &point{}

	setPoint(points)

	printText("x = ")
	printInt1(points.x)
	printInt2(points.x)
	printText(", y = ")
	printInt1(points.y)
	printInt2(points.y)
	printText("\n")
}

package piscine

type food struct {
	burger  int
	chips   int
	nuggets int
}

func FoodDeliveryTime(order string) int {
	time := food{}
	time.burger = 15
	time.chips = 10
	time.nuggets = 12

	if order == "burger" {
		return time.burger
	} else if order == "chips" {
		return time.chips
	} else if order == "nuggets" {
		return time.nuggets
	} else {
		return 404
	}
}

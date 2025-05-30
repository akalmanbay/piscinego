package piscine

func ConcatParams(args []string) string {
	res := ""
	for i := range args {
		res = res + args[i]
		if i != len(args)-1 {
			res = res + "\n"
		}
	}
	return res
}

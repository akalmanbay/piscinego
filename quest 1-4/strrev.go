package piscine

func StrRev(s string) string {
	res := []byte(s)
	cnt := len(s) - 1
	for i, val := range s {
		res[cnt-i] = byte(val)
	}
	return string(res)
}

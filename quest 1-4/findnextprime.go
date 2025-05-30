package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	} else if nb == 3 {
		return 3
	} else if nb%2 == 0 {
		nb = nb + 1
	}

	for nextp := nb; nextp < 9223372036854775807; {
		for i := 3; i < nextp; {
			if nextp%i == 0 {
				break
			} else if i+2 == nextp {
				return nextp
			}
			i = i + 2
		}
		nextp = nextp + 2
	}
	return 9223372036854775807
}

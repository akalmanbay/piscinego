package piscine

func PodiumPosition(podium [][]string) [][]string {
	j := 0
	for i := len(podium) - 1; i >= 0; i-- {
		podium[j], podium[i] = podium[i], podium[j]
		j++
	}
	return podium
}

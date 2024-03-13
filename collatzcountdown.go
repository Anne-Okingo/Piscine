package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	steps := 0
	for start != 1 {
		if start%2 != 0 {
			start = start*3 + 1
		} else {
			start = start / 2
		}
		steps++
	}
	return steps
}

package numbersutil

func ForcePositiveValue(value int) int {
	if value <= 0 {
		return 1
	}
	return value
}

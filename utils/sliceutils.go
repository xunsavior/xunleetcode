package utils

//Sum ...
func Sum(nums []float64) float64 {
	sum := 0.0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// StringToRunes ...
func StringToRunes(str string) []rune {
	runes := make([]rune, 0, len(str))
	for _, v := range str {
		runes = append(runes, v)
	}
	return runes
}

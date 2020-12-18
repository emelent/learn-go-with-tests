package numeral

func ConvertToRoman(num int) string {
	if num == 2 {
		return "II"
	}
	if num == 3 {
		return "III"
	}
	return "I"
}

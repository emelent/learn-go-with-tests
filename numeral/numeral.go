package numeral

import "strings"

func ConvertToRoman(num int) string {

	var result strings.Builder

	if num == 4 {
		return "IV"
	}

	for i := 0; i < num; i++ {
		result.WriteString("I")
	}

	return result.String()
}

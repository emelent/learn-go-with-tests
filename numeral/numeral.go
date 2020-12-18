package numeral

import "strings"

func ConvertToRoman(num int) string {

	var result strings.Builder

	for i := 0; i < num; i++ {
		result.WriteString("I")
	}

	return result.String()
}

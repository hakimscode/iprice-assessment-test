package string_manipulation

import "strings"

func GetUppercase(str string) string {
	return strings.ToUpper(str)
}

func GetAlternateCase(str string) string {
	strs := strings.Split(str, "")

	result := ""
	upper := true
	for _, str := range strs {
		if upper {
			result += strings.ToUpper(str)
			upper = false
		} else {
			result += strings.ToLower(str)
			upper = true
		}
	}

	return result
}

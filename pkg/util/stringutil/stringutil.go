package stringutil

import (
	"unicode/utf8"

	"github.com/asaskevich/govalidator"
)

func Diff(Base, Exclude []string) []string {
	ExcludeMap := make(map[string]bool, 0)
	result := make([]string, 0)

	for _, exc := range Exclude {
		ExcludeMap[exc] = true
	}

	for _, base := range Base {
		if !ExcludeMap[base] {
			result = append(result, base)
		}
	}

	return result
}

func Unique(ss []string) []string {
	uniqueMap := make(map[string]bool, 0)
	result := make([]string, 0)

	for _, s := range ss {
		uniqueMap[s] = true
	}

	for key := range uniqueMap {
		result = append(result, key)
	}

	return result
}

func CamelCaseToUnderscore(s string) string {
	return govalidator.CamelCaseToUnderscore(s)
}

func UnderscoreToCamelCase(s string) string {
	return govalidator.UnderscoreToCamelCase(s)
}

func FindString(find string, ss []string) int {
	for i, s := range ss {
		if find == s {
			return i
		}
	}

	return -1
}

func StringIn(find string, ss []string) bool {
	return FindString(find, ss) != -1
}

func Reverse(str string) string {
	n := len(str)
	result := make([]byte, n)
	start := 0

	for start < n {
		r, width := utf8.DecodeRuneInString(str[start:])
		utf8.EncodeRune(result[n-width-start:], r)
		start += width
	}

	return string(result)
}

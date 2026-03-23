package text

import (
	"fmt"
	"regexp"
)

func ReplacePlaceholders(str string, params []any) string {
	pattern := regexp.MustCompile(`\$\d+`)

	for _, v := range params {
		loc := pattern.FindStringIndex(str)
		if loc == nil {
			break
		}

		str = fmt.Sprintf("%v%v%v", str[:loc[0]], v, str[loc[1]:])
	}

	return str
}

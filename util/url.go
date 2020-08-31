package util

import (
	"regexp"
)

var (
	REGEXP_EXT = regexp.MustCompile("\\.([a-zA-Z0-9]+)\\?")
)

func GetExt(url string) string {
	result := REGEXP_EXT.FindStringSubmatch(url)
	if len(result) == 2 {
		return result[1]
	}
	return ""
}

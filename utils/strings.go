package utils

import (
	"fmt"
	"github.com/gosimple/slug"
	"regexp"
)

func SlugString(s string) (string, error) {
	slugged := slug.Make(s)
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}
	clean := reg.ReplaceAllString(slugged, "")
	return clean, nil
}

func ConcatStrings(s1 string, s2 string) string {
	return fmt.Sprintf("%s%s", s1, s2)
}

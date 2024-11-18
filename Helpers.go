package main

import "strings"

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func cleanStringForSearch(s string) string {
	s = strings.Replace(s, "&", "", -1)
	s = strings.Replace(s, "/", "", -1)
	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	s = strings.Replace(s, "[", "", -1)
	s = strings.Replace(s, "]", "", -1)
	s = strings.Replace(s, "-", "", -1)
	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	s = strings.Replace(s, "feat.", "", -1)
	s = strings.Replace(s, "ft.", "", -1)

	return standardizeSpaces(s)
}

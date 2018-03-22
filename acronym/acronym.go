// Package acronym provides the acronym to a string input
package acronym

import (
	"log"
	"regexp"
	"strings"
)

// Abbreviate gives the acronym of the string s
func Abbreviate(s string) string {
	var abv string // the abbreviation of s

	// Remove any non letters
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	noPunc := reg.ReplaceAllString(s, " ")

	// Loop over each word and add the first letter to abv
	for _, v := range strings.Fields(noPunc) {
		abv = abv + strings.ToUpper(string(v[0]))
	}

	return abv
}

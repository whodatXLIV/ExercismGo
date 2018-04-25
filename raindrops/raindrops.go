//Package raindrops converts a number to a string
package raindrops

import (
	"strconv"
)

//Convert takes an integer and outputs a string
func Convert(n int) string {
	s := ""
	n3 := n%3 == 0
	n5 := n%5 == 0
	n7 := n%7 == 0

	if !(n3 || n5 || n7) {
		return strconv.Itoa(n)
	} else {

		if n3 {
			s = s + "Pling"
		}
		if n5 {
			s = s + "Plang"
		}
		if n7 {
			s = s + "Plong"
		}
		return s
	}

}

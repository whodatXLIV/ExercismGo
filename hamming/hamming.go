// Package hamming calculates the difference between two DNA strands
package hamming

import (
	"errors"
)

// Distance takes DNA strands as arguments a and b and returns their Hamming distance
func Distance(a, b string) (int, error) {
	cnt := 0
	if a == b {
		return cnt, nil
	} else if len(a) != len(b) {
		return -1, errors.New("DNA string lengths not equal")
	} else {
		for i := range a {
			if a[i] != b[i] {
				cnt += 1
			}
		}
		return cnt, nil
	}
}

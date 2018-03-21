// Package bob provides Bob's responses to conversation
package bob

import "strings"

// Hey returns a string of Bob's response to remark
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)
	if len(remark) == 0 {
		//Remark is silence
		return "Fine. Be that way!"
	}

	upperRemark := strings.ToUpper(remark)
	isWords := strings.ContainsAny(upperRemark, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	isQuestion := strings.HasSuffix(remark, "?")
	isYell := remark == upperRemark && isWords

	switch {
	case isQuestion && isYell: //Remark is a yelled question
		return "Calm down, I know what I'm doing!"

	case isQuestion: // Remark is a question
		return "Sure."

	case isYell: //Remark is yelled
		return "Whoa, chill out!"

	default: // default
		return "Whatever."
	}
}

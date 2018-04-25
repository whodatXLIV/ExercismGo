// Package twofer should have a package comment that summarizes what it's about.
package twofer

// ShareWith should have a comment documenting it.
func ShareWith(s string) string {
	if len(s) > 0 {
		return "One for " + s + ", one for me."
	} else {
		return "One for you, one for me."
	}
}

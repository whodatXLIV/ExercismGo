// Package gigasecond is used to add 1 gigasecond to a given time
package gigasecond

import "time"

// AddGigasecond returns the sum of input time t and 1 gigasecond
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(1e9 * time.Second)
	return t
}

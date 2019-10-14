// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

import "time"

// AddGigasecond add one gigasenco to a given time
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(1000000000 * time.Second)
	return t
}

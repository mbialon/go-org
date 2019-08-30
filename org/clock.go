package org

import (
	"fmt"
	"regexp"
)

type Clock struct {
	Start    string
	End      string
	Duration string
}

func (c Clock) String() string {
	return fmt.Sprintf("CLOCK, start: %s, end: %s, duration: %s", c.Start, c.End, c.Duration)
}

var clockRegexp = regexp.MustCompile(`^(\s*)CLOCK: (.*)`)

func lexClock(line string) (token, bool) {
	if m := clockRegexp.FindStringSubmatch(line); m != nil {
		return token{"clock", len(m[1]), m[2], m}, true
	}
	return nilToken, false
}

func (d *Document) parseClock(i int, parentStop stopFn) (int, Node) {
	//_ := d.tokens[i].content
	return 1, Clock{}
}

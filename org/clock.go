package org

import (
	"fmt"
	"regexp"
	"time"
)

type Clock struct {
	Start    time.Time
	End      time.Time
	Duration time.Duration
}

func (c Clock) String() string {
	return fmt.Sprintf("CLOCK, start: %s, end: %s, duration: %s", c.Start, c.End, c.Duration)
}

var (
	clockRegexp        = regexp.MustCompile(`^(\s*)CLOCK: (.*)`)
	clockContentRegexp = regexp.MustCompile(`\[(.+)]--\[(.+)]\s*=>(.+)`)
)

func lexClock(line string) (token, bool) {
	if m := clockRegexp.FindStringSubmatch(line); m != nil {
		return token{"clock", len(m[1]), m[2], m}, true
	}
	return nilToken, false
}

func (d *Document) parseClock(i int, parentStop stopFn) (int, Node) {
	m := clockContentRegexp.FindStringSubmatch(d.tokens[i].content)
	if m == nil {
		return 0, nil
	}
	_, start := d.parseTimestamp(m[1], 0)
	if start == nil {
		return 0, nil
	}
	_, end := d.parseTimestamp(m[2], 0)
	if end == nil {
		return 0, nil
	}
	return 1, Clock{
		Start: start.(Timestamp).Time,
		End:   end.(Timestamp).Time,
	}
}

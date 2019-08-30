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
	clockRegexp        = regexp.MustCompile(`^(\s*)CLOCK: (.*)$`)
	clockContentRegexp = regexp.MustCompile(`\[(.+)]--\[(.+)]\s*=>(.+)`)
)

func lexClock(line string) (token, bool) {
	if m := clockRegexp.FindStringSubmatch(line); m != nil {
		return token{"clock", len(m[1]), m[2], m}, true
	}
	return nilToken, false
}

func (d *Document) parseClock(i int, parentStop stopFn) (int, Node) {
	s := d.tokens[i].content
	d.Log.Printf("CLOCK, parse: %s", s)
	m := clockContentRegexp.FindStringSubmatch(s)
	d.Log.Printf("CLOCK, matches: %v", m)
	if m == nil {
		return 0, nil
	}
	d.Log.Printf("CLOCK, parse T1: %s", m[1])
	_, start := d.parseTimestamp("<"+m[1]+">", 0)
	d.Log.Printf("CLOCK, T1: %v", start)
	if start == nil {
		return 0, nil
	}
	d.Log.Printf("CLOCK, parse T2: %s", m[2])
	_, end := d.parseTimestamp("<"+m[2]+">", 0)
	d.Log.Printf("CLOCK, T2: %v", end)
	if end == nil {
		return 0, nil
	}
	return 1, Clock{
		Start: start.(Timestamp).Time,
		End:   end.(Timestamp).Time,
	}
}

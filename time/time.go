package time

import "time"

func Between(t, min, max time.Time) bool {
	return t.After(min) && t.Before(max)
}

func TimeIn(tz string, t time.Time) time.Time {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return t
	}
	return t.In(loc)
}

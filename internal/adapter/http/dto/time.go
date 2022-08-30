package dto

import "time"

func toTime(t string) (time.Duration, error) {
	return time.ParseDuration(t)
}

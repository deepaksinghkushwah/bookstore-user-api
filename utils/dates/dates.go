package dates

import "time"

const apiDateLayout = "2006-01-02 15:04:05"

// GetNow return date litral
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString return current time string
func GetNowString() string {
	now := GetNow().UTC()
	return now.Format(apiDateLayout)
}

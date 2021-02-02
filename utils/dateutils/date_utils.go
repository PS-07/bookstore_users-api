package dateutils

import "time"

const apiDateLayout = "2006-01-02T15:04:05Z"

// GetNow func
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString func
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

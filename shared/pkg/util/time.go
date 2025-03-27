package util

import "time"

func DateStringToTime(dateString string) (time.Time, error) {
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return time.Now(), err
	}
	date, err := time.ParseInLocation("2006-01-02 15:04:05", dateString, location)
	if err != nil {
		return time.Now(), err
	}
	return date, nil
}

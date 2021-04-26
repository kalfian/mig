package lib

import (
	"strings"
	"time"
)

// Parse from RFC3339Nano format to 2006-01-02 15:04:05
func ParseRFC3339Nano(date string) string {
	dt, err := time.Parse(time.RFC3339Nano, date)
	if err != nil {
		return date
	}
	dateParsed := dt.Format("2006-01-02 15:04:05")
	return dateParsed
}

// Parse from 2006-01-02 to 2006-01-02 15:04:05
func ParseISO(date string) string {
	dt, err := time.Parse("2006-01-02", date)
	if err != nil {
		return date
	}
	dateParsed := dt.Format("2006-01-02 15:04:05")
	return dateParsed
}

// Parse from 2006-01-02 to 02 Jan 2006
func Parseto02Jan2006(date string) string {
	dt, err := time.Parse("2006-01-02", date)
	if err != nil {
		return date
	}
	dateParsed := dt.Format("02 Jan 2006")
	return dateParsed
}

// Check and fill second of date with format 2006-01-02 15:04:05
func FillISOSeconds(date string) string {
	dt := strings.Split(date, ":")
	if dt[2] == "00" {
		dt[2] = time.Now().Format("05")
		return strings.Join(dt, ":")
	}
	return date
}

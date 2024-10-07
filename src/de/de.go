package de

import (
	"fmt"
	"regexp"
	"simonneutert/strava-export-organizer/src/helpers"
	"strings"
	"time"
)

func MapDateDe(datestr string) string {
	if regexp.MustCompile(`\d{2}\.\d{2}\.\d{4}`).MatchString(datestr) == false {
		return "1986-01-01T15:04:05.999999999Z"
	}
	day := strings.Split(datestr, ".")[0]
	month := strings.Split(datestr, ".")[1]
	year := strings.Split(datestr, ".")[2][:4]
	return fmt.Sprintf("%s-%s-%sT15:04:05.999999999Z", year, month, day)
}

func MapFilenameDe(filename string, datestr string) string {
	dateParsed, err := time.Parse("2006-01-02T15:04:05.999999999Z", datestr)
	if err != nil {
		dateParsed, _ = time.Parse("2006-01-02T15:04:05.999999999Z", "1986-01-01T15:04:05.999999999Z")
	}
	fmt.Println(dateParsed)
	year, month, day := dateParsed.Date()

	activityFilename := helpers.FilenameFromPath(filename)
	return fmt.Sprintf("%d%02d%02d-%s", year, month, day, activityFilename)
}

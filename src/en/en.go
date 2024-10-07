package en

import (
	"fmt"
	"simonneutert/strava-export-organizer/src/helpers"
	"time"
)

func MapDateEn(datestr string) string {
	t, err := time.Parse("Jan 2, 2006, 3:04:05 PM", datestr)
	if err != nil {
		return "1986-01-01T15:04:05.999999999Z"
	}
	date := t.Format("2006-01-02")
	return fmt.Sprintf("%sT15:04:05.999999999Z", date)
}

func MapFilenameEn(filename string, datestr string) string {

	dateParsed, err := time.Parse("2006-01-02T15:04:05.999999999Z", datestr)
	if err != nil {
		dateParsed, _ = time.Parse("2006-01-02T15:04:05.999999999Z", "1986-01-01T15:04:05.999999999Z")
	}

	year, month, day := dateParsed.Date()
	activityFilename := helpers.FilenameFromPath(filename)
	return fmt.Sprintf("%d%02d%02d-%s", year, month, day, activityFilename)
}

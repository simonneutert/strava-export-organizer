package main

import (
	"testing"
	"time"
)

func TestTimeParsingAndFormatting(t *testing.T) {
	record := "Sep 26, 2016, 6:37:51 AM" // Define record for the sake of example
	parsedTime, _ := time.Parse("Jan 2, 2006, 3:04:05 PM", record)
	day := parsedTime.Format("2006-01-02")
	if day != "2016-09-26" {
		t.Errorf("Expected 2016-09-26, got %s", day)
	}
}

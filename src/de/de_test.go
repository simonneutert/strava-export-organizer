package de_test

import (
	"simonneutert/strava-export-organizer/src/de"
	"testing"
)

func TestMapDateDe(t *testing.T) {
	tests := []struct {
		name     string
		datestr  string
		expected string
	}{
		{"Valid date format", "02.01.2006", "2006-01-02T15:04:05.999999999Z"},
		{"Invalid date format", "not-a-date", "1986-01-01T15:04:05.999999999Z"},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := de.MapDateDe(tt.datestr)
			if result != tt.expected {
				t.Errorf("MapDateDe(%s) = %s; want %s", tt.datestr, result, tt.expected)
			}
		})
	}
}

func TestMapFilenameDe(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		datestr  string
		expected string
	}{
		{"Valid filename and date", "abc.gpx", "2006-01-02T15:04:05.999999999Z", "20060102-abc.gpx"},
		{"Invalid date in filename", "def.fit", "not-a-date", "19860101-def.fit"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := de.MapFilenameDe(tt.filename, tt.datestr)
			if result != tt.expected {
				t.Errorf("MapFilenameDe(%s, %s) = %s; want %s", tt.filename, tt.datestr, result, tt.expected)
			}
		})
	}
}

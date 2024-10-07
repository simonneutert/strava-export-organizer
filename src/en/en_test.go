package en_test

import (
	"simonneutert/strava-export-organizer/src/en"
	"testing"
)

func TestMapDateEn(t *testing.T) {
	tests := []struct {
		name     string
		datestr  string
		expected string
	}{
		{"Valid date format", "Jan 2, 2006, 3:04:05 PM", "2006-01-02T15:04:05.999999999Z"},
		{"Invalid date format", "not-a-date", "1986-01-01T15:04:05.999999999Z"},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := en.MapDateEn(tt.datestr)
			if result != tt.expected {
				t.Errorf("MapDateEn(%s) = %s; want %s", tt.datestr, result, tt.expected)
			}
		})
	}
}

func TestMapFilenameEn(t *testing.T) {
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
			result := en.MapFilenameEn(tt.filename, tt.datestr)
			if result != tt.expected {
				t.Errorf("MapFilenameEn(%s, %s) = %s; want %s", tt.filename, tt.datestr, result, tt.expected)
			}
		})
	}
}

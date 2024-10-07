package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"os/exec"
	"simonneutert/strava-export-organizer/src/de"
	"simonneutert/strava-export-organizer/src/en"
	"strings"
)

// define a global struct to hold language configuration
var languageConfig map[string]map[string]string

type dateMapper func(string) string
type filenameMapper func(string, string) string

func init() {
	languageConfig = make(map[string]map[string]string)

	languageConfig["en"] = map[string]string{
		"filename":     "Filename",
		"activityType": "Activity Type",
		"activityDate": "Activity Date",
		"activityName": "Activity Name",
	}

	languageConfig["de"] = map[string]string{
		"filename":     "Dateiname",
		"activityType": "Aktivitätsart",
		"activityDate": "Aktivitätsdatum",
		"activityName": "Name der Aktivität",
	}
}

func main() {
	targetDir := "./export_mapped"
	language := "de"

	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Usage: go run main.go [language]")
		os.Exit(1)
	}

	if os.Args[1] != "en" && os.Args[1] != "de" {
		fmt.Println("Language must be 'en' or 'de'")
		os.Exit(1)
	}

	language = os.Args[1]

	csvFile, _ := os.Open(fmt.Sprintf("activities.csv"))
	reader := csv.NewReader(csvFile)
	reader.Comma = ','
	reader.LazyQuotes = true
	records, _ := reader.ReadAll()

	mappedData := mapData(records, language)
	moveFiles(language, mappedData, targetDir)
	gunzipFiles()
}

func moveFiles(language string, mappedData []map[string]string, targetDir string) {
	activityDateKey := languageConfig[language]["activityDate"]
	activityTypeKey := languageConfig[language]["activityType"]
	activityFilename := languageConfig[language]["filename"]

	for _, row := range mappedData {
		if row[activityFilename] == "" {
			continue
		}

		year := strings.Split(row[activityDateKey], "-")[0]
		activityType := row[activityTypeKey]

		activityYearDir := fmt.Sprintf("%s_%s", activityType, year)
		newDir := fmt.Sprintf("%s/%s/%s", targetDir, activityType, activityYearDir)
		os.MkdirAll(newDir, os.ModePerm)

		oldPath := fmt.Sprintf("./%s", row[activityFilename])
		newPath := fmt.Sprintf("%s/%s", newDir, row["Filenamenew"])
		copyFile(oldPath, newPath)
		fmt.Println(newPath)
	}
	fmt.Println("\n\nDone 🚀")
	return
}

func mapData(records [][]string, language string) []map[string]string {
	var result []map[string]string
	if language == "de" {
		result = mapGenericRows(records, language, de.MapDateDe, de.MapFilenameDe)
	}
	if language == "en" {
		result = mapGenericRows(records, language, en.MapDateEn, en.MapFilenameEn)
	}
	return result
}

func mapGenericRows(records [][]string, language string, dateMapper dateMapper, filenameMapper filenameMapper) []map[string]string {
	var mappedData []map[string]string
	for _, record := range records[1:] {
		row := make(map[string]string)
		row[languageConfig[language]["filename"]] = record[12]
		row[languageConfig[language]["activityType"]] = record[3]
		row[languageConfig[language]["activityName"]] = strings.ReplaceAll(record[2], "[^0-9A-Za-z]", "_")
		row[languageConfig[language]["activityDate"]] = dateMapper(record[1])
		row["Filenamenew"] = filenameMapper(record[12], row[languageConfig[language]["activityDate"]])
		mappedData = append(mappedData, row)
	}
	return mappedData
}

func copyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func gunzipFiles() {
	exec.Command("/bin/sh", "-c", "gunzip -r ./export_mapped").Output()
}

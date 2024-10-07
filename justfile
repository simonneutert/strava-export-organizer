build:
  rm -rf dist; rm -rf export_mapped
  go test ./... && goreleaser release --snapshot --clean 

test language:
  @echo "Running tests for {{language}}"
  @echo "Please remember to clean `export_1/export_mapped` yourself 🚀"
  go build
  go test ./...
  cp strava-export-organizer export_1/strava-export-organizer
  cd export_1 && ./strava-export-organizer {{language}}
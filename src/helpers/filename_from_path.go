package helpers

import "strings"

func FilenameFromPath(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

package util

import (
	"fmt"
	"runtime"
	"strings"
)

// Get current file's absolute path
func GetBasePath() (string, error) {
	_, filePath, _, ok := runtime.Caller(1) // Use caller(1) to get the path of the function that calls this helper
	if !ok {
		return "", fmt.Errorf("could not get caller information")
	}

	return filePath, nil
}

// surrounds strings with quotes, then joins them with a comma and space
func JoinStrings(strs []string) string {

	quotedStrs := []string{}
	for _, str := range strs {
		quotedStrs = append(quotedStrs, "\""+str+"\"")
	}

	return strings.Join(quotedStrs, ", ")
}

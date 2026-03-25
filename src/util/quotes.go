package util

import "fmt"

func WrapQuotes(s string) string {
	return fmt.Sprintf("\"%s\"", s)
}

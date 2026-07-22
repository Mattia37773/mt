package base

import "regexp"

// Removes ansi color codes
func StripANSI(str string) string {
	ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	return ansiRegex.ReplaceAllString(str, "")
}

//
// exitFunc = func(code int) {
// exited = true
//     }

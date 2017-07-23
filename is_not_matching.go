package validator

import "regexp"

// IsNotMatching matches the given s to the given regex and returns a true if
// it does not match and false if it does.
func IsNotMatching(s, regex, msg string) (bool, string) {
	b := !regexp.MustCompile(regex).MatchString(s)
	if b == false {
		if msg == "" {
			msg = "String does not match regex"
		}
	}
	return b, msg
}

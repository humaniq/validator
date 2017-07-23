package validator

// IsStrEmpty check string's length (including multi byte strings)
func IsStrEmpty(str, msg string) (bool, string) {
	if len(str) > 0 {
		if msg == "" {
			msg = "String length is greater than 0"
		}
		return false, msg
	}

	return true, msg
}

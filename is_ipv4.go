package validator

import (
	"net"
	"strings"
)

// IsIPv4 check if the string is an IP version 4.
func IsIPv4(str, msg string) (bool, string) {
	b := (net.ParseIP(str) != nil && strings.Contains(str, "."))
	if b == false {
		if msg == "" {
			msg = "String is not of IPv4 format"
		}
	}
	return b, msg
}

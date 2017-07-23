package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNotMatching(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		str string
		reg string
		msg string
		e   expected
	}{
		{str: "1928-06-05", reg: `[12][0-9]{3}-[01][0-9]-[0-3][0-9]`, msg: ``, e: expected{b: false, msg: "String does not match regex"}},
		{str: "1928/03/02", reg: `[12][0-9]{3}-[01][0-9]-[0-3][0-9]`, msg: ``, e: expected{b: true, msg: ""}},
	}
	for _, test := range tests {
		b, msg := IsNotMatching(test.str, test.reg, test.msg)
		assert.Equal(t, b, test.e.b)
		assert.Equal(t, msg, test.e.msg)
	}
}

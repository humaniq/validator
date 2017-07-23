package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStrEmpty(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param string
		msg   string
		e     expected
	}{
		{param: "1", msg: "", e: expected{b: false, msg: "String length is greater than 0"}},
		{param: "", msg: "zoo", e: expected{b: true, msg: "zoo"}},
	}

	for _, test := range tests {
		b, msg := IsStrEmpty(test.param, test.msg)
		assert.Equal(t, b, test.e.b)
		assert.Equal(t, msg, test.e.msg)
	}
}

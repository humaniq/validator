package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIPv4(t *testing.T) {
	t.Parallel()

	tests := []struct {
		param string
		msg   string
		e     expected
	}{
		{param: "foo", msg: "", e: expected{b: false, msg: "String is not of IPv4 format"}},
		{param: "127.0.0.1", msg: "foo", e: expected{b: true, msg: "foo"}},
		{param: "0.0.0.0", msg: "foo", e: expected{b: true, msg: "foo"}},
		{param: "256.255.255.255", msg: "foo", e: expected{b: false, msg: "foo"}},
		{param: "1.2.3.4", msg: "foo", e: expected{b: true, msg: "foo"}},
		{param: "::1", msg: "foo", e: expected{b: false, msg: "foo"}},
		{param: "2001:db8:0000:1:1:1:1:1", msg: "foo", e: expected{b: false, msg: "foo"}},
		{param: "300.0.0.0", msg: "foo", e: expected{b: false, msg: "foo"}},
	}

	for _, test := range tests {
		b, msg := IsIPv4(test.param, test.msg)
		assert.Equal(t, b, test.e.b)
		assert.Equal(t, msg, test.e.msg)
	}
}

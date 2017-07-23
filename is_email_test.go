package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmail(t *testing.T) {
	t.Parallel()

	defaultMsg := "String is formatted to be an email address"
	var tests = []struct {
		param string
		msg   string
		e     expected
	}{
		{param: ``, msg: ``, e: expected{b: false, msg: defaultMsg}},
		{param: `foo@bar.com`, msg: `valid email`, e: expected{b: true, msg: `valid email`}},
		{param: `x@bx.com`, msg: `valid email`, e: expected{b: true, msg: `valid email`}},
		{param: `foo@bar.com.au`, msg: `valid email`, e: expected{b: true, msg: `valid email`}},
		{param: `foo+bar@bar.com`, msg: `valid email`, e: expected{b: true, msg: `valid email`}},
		{param: `foo@bar.coffee`, msg: `valid email`, e: expected{b: true, msg: `valid email`}},
		{param: `foo@bar.中文网`, msg: `valid email`, e: expected{b: true, msg: `valid email`}},
		// {`invalidemail@`, false},
		// {`invalid.com`, false},
		// {`@invalid.com`, false},
		// {`test|123@m端ller.com`, true},
		// {`hans@m端ller.com`, true},
		// {`hans.m端ller@test.com`, true},
		// {`foo.Babbie@DomaIn.cOM`, true},
		// {`manbearpig@DOMAIN.CO.UK`, true},
		// {`very.(),:;<>[]".VERY."very@\ "very".unusual@strange.example.com`, true},
	}
	for _, test := range tests {
		b, msg := IsEmail(test.param, test.msg)
		assert.Equal(t, b, test.e.b)
		assert.Equal(t, msg, test.e.msg)
	}
}

package codes_test

import (
	"fmt"
	"testing"

	"github.com/creikey/rpgpt/server/codes"
)

func TestParseUserCode(t *testing.T) {
	for _, tt := range []struct {
		input string
		code  codes.UserCode
		err   error
	}{
		{
			input: "AAAA",
			code:  codes.UserCode(0),
			err:   nil,
		},
		{
			input: "AAAB",
			code:  codes.UserCode(1),
			err:   nil,
		},
		{
			input: "BAAA",
			code:  codes.UserCode(46656),
			err:   nil,
		},
	} {
		t.Run(tt.input, func(t *testing.T) {
			code, err := codes.ParseUserCode(tt.input)
			if got, want := err, tt.err; got != want {
				t.Fatalf("err=%v, want=%v", got, want)
			}
			if got, want := code, tt.code; got != want {
				t.Errorf("code=%v, want=%v", got, want)
			}
		})
	}
}

func TestCodeToString(t *testing.T) {
	for _, tt := range []struct {
		input   codes.UserCode
		encoded string
		err     error
	}{
		{
			input:   codes.UserCode(1),
			encoded: "AAAB",
			err:     nil,
		},
	} {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			encoded, err := codes.CodeToString(tt.input)
			if got, want := err, tt.err; got != want {
				t.Fatalf("err=%v, want=%v", got, want)
			}
			if got, want := encoded, tt.encoded; got != want {
				t.Errorf("encoded=%v, want=%v", got, want)
			}
		})
	}
}

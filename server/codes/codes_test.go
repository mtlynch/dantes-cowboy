package codes_test

import (
	"fmt"
	"testing"

	"github.com/creikey/rpgpt/server/codes"
)

func TestParseUserCode(t *testing.T) {
	for _, tt := range []struct {
		input   string
		code    codes.UserCode
		errText string
	}{
		{
			input:   "AAAA",
			code:    codes.UserCode(0),
			errText: "",
		},
		{
			input:   "AAAB",
			code:    codes.UserCode(1),
			errText: "",
		},
		{
			input:   "BAAA",
			code:    codes.UserCode(46656),
			errText: "",
		},
		{
			input:   "9999",
			code:    codes.UserCode(1679615),
			errText: "",
		},
		{
			input:   "AAAa",
			code:    codes.UserCode(0),
			errText: "failed to find place's number AAAa",
		},
		{
			input:   "",
			code:    codes.UserCode(0),
			errText: "string to deconvert is not of length 4: ",
		},
		{
			input:   "AAA",
			code:    codes.UserCode(0),
			errText: "string to deconvert is not of length 4: AAA",
		},
		{
			input:   "AAAAA",
			code:    codes.UserCode(0),
			errText: "string to deconvert is not of length 4: AAAAA",
		},
	} {
		t.Run(tt.input, func(t *testing.T) {
			code, err := codes.ParseUserCode(tt.input)
			if got, want := errToString(err), tt.errText; got != want {
				t.Fatalf("err=%v, want=%v", got, want)
			}
			if err != nil {
				return
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
		errText string
	}{
		{
			input:   codes.MinUserCode,
			encoded: "AAAA",
			errText: "",
		},
		{
			input:   codes.UserCode(1),
			encoded: "AAAB",
			errText: "",
		},
		{
			input:   codes.UserCode(46656),
			encoded: "BAAA",
			errText: "",
		},
		{
			input:   codes.MaxUserCode,
			encoded: "9999",
			errText: "",
		},
	} {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			encoded, err := codes.CodeToString(tt.input)
			if got, want := errToString(err), tt.errText; got != want {
				t.Fatalf("err=%v, want=%v", got, want)
			}
			if got, want := encoded, tt.encoded; got != want {
				t.Errorf("encoded=%v, want=%v", got, want)
			}
		})
	}
}

func errToString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

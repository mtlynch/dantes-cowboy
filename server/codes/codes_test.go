package codes_test

import (
	"fmt"
	"testing"

	"github.com/creikey/rpgpt/server/codes"
)

func TestFromString(t *testing.T) {
	for _, tt := range []struct {
		input   string
		code    codes.UserCode
		errText string
	}{
		{
			input:   "AAAA",
			code:    mustMakeUserCode(0),
			errText: "",
		},
		{
			input:   "AAAB",
			code:    mustMakeUserCode(1),
			errText: "",
		},
		{
			input:   "BAAA",
			code:    mustMakeUserCode(46656),
			errText: "",
		},
		{
			input:   "9999",
			code:    mustMakeUserCode(1679615),
			errText: "",
		},
		{
			input:   "AAAa",
			code:    codes.UserCode{},
			errText: "failed to find place's number AAAa",
		},
		{
			input:   "",
			code:    codes.UserCode{},
			errText: "string to deconvert is not of length 4: ",
		},
		{
			input:   "AAA",
			code:    codes.UserCode{},
			errText: "string to deconvert is not of length 4: AAA",
		},
		{
			input:   "AAAAA",
			code:    codes.UserCode{},
			errText: "string to deconvert is not of length 4: AAAAA",
		},
	} {
		t.Run(tt.input, func(t *testing.T) {
			code, err := codes.FromString(tt.input)
			if got, want := errToString(err), tt.errText; got != want {
				t.Fatalf("err=%v, want=%v", got, want)
			}
			if err != nil {
				return
			}
			if got, want := code, tt.code; got != want {
				t.Errorf("code=%v, want=%v", got.String(), want.String())
			}
		})
	}
}

func TestFromInt(t *testing.T) {
	for _, tt := range []struct {
		input   int
		code    codes.UserCode
		errText string
	}{
		{
			input:   0,
			code:    mustMakeUserCode(0),
			errText: "",
		},
		{
			input:   codes.MaxUserCode.Int(),
			code:    mustMakeUserCode(1679615),
			errText: "",
		},
		{
			input:   codes.MinUserCode.Int() - 1,
			code:    codes.UserCode{},
			errText: "user code -1 is out of range, must be between 0 and 1679615",
		},
		{
			input:   codes.MaxUserCode.Int() + 1,
			code:    codes.UserCode{},
			errText: "user code 1679616 is out of range, must be between 0 and 1679615",
		},
	} {
		t.Run(fmt.Sprintf("%d", tt.input), func(t *testing.T) {
			code, err := codes.FromInt(tt.input)
			if got, want := errToString(err), tt.errText; got != want {
				t.Fatalf("err=%v, want=%v", got, want)
			}
			if err != nil {
				return
			}
			if got, want := code, tt.code; got != want {
				t.Errorf("code=%v, want=%v", got.String(), want.String())
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
			input:   mustMakeUserCode(1),
			encoded: "AAAB",
			errText: "",
		},
		{
			input:   mustMakeUserCode(46656),
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
			if got, want := tt.input.String(), tt.encoded; got != want {
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

func mustMakeUserCode(val int) codes.UserCode {
	uc, err := codes.FromInt(val)
	if err != nil {
		panic(err)
	}
	return uc
}

package codes

import (
	"fmt"
	"math/rand"
)

const (
	// A-Z and 0-9, four digits means this many codes
	totalPossibleCodes = 36 * 36 * 36 * 36

	CodeStringLength = 4
)

type UserCode struct {
	inner int
}

type ErrInvalidCodeStringLength struct {
	Input string
}

func (e ErrInvalidCodeStringLength) Error() string {
	return fmt.Sprintf("string to deconvert is not of length %d: %s", CodeStringLength, e.Input)
}

type ErrInvalidCodeStringCharacter struct {
	Input string
}

func (e ErrInvalidCodeStringCharacter) Error() string {
	return fmt.Sprintf("failed to find place's number %s", e.Input)
}

type ErrCodeOutOfRange struct {
	Input int
}

func (e ErrCodeOutOfRange) Error() string {
	return fmt.Sprintf("user code %d is out of range, must be between %d and %d", e.Input, MinUserCode.inner, MaxUserCode.inner)
}

func intPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

var (
	numberToChar = [...]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	MinUserCode = UserCode{0}
	MaxUserCode = UserCode{totalPossibleCodes - 1}
)

func Random() UserCode {
	return UserCode{rand.Intn(totalPossibleCodes)}
}

func FromInt(val int) (UserCode, error) {
	if val < MinUserCode.inner || val > MaxUserCode.inner {
		return UserCode{}, ErrCodeOutOfRange{val}
	}
	return UserCode{val}, nil
}

func FromString(s string) (UserCode, error) {
	asRune := []rune(s)
	if len(asRune) != CodeStringLength {
		return UserCode{}, ErrInvalidCodeStringLength{Input: s}
	}
	var val int
	for place := CodeStringLength - 1; place >= 0; place-- {
		index := CodeStringLength - 1 - place
		curDigitNum := 0
		found := false
		for i, letter := range numberToChar {
			if letter == asRune[index] {
				curDigitNum = i
				found = true
			}
		}
		if !found {
			return UserCode{}, ErrInvalidCodeStringCharacter{Input: s}
		}
		val += curDigitNum * intPow(36, place)
	}
	return UserCode{val}, nil
}

func (uc UserCode) String() string {
	toReturn := [CodeStringLength]rune{'A', 'A', 'A', 'A'}

	value := uc.inner

	for place := CodeStringLength - 1; place >= 0; place-- {
		index := CodeStringLength - 1 - place
		currentPlaceValue := value / intPow(36, place)
		value -= currentPlaceValue * intPow(36, place)
		toReturn[index] = numberToChar[currentPlaceValue]
	}

	return string(toReturn[:])
}

func (uc UserCode) Int() int {
	return uc.inner
}

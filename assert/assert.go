package assert

import (
	"fmt"

	"github.com/muxit-studio/color"
	"github.com/muxit-studio/columnize"
)

type T interface {
	Errorf(format string, args ...interface{})
}

func Equal(t T, got interface{}, expected interface{}, message string) bool {
	return assert(t, got, expected, message, true)
}

func NotEqual(t T, got interface{}, expected interface{}, message string) bool {
	return assert(t, got, expected, message, false)
}

func assert(t T, got interface{}, expected interface{}, message string, expectation bool) bool {
	errorMessage := []string{
		"\n",
		fmt.Sprintf(color.BWhite("message: | \"%v\""), message),
		"\n",
		fmt.Sprintf(color.BGreen("expected: | %v"), expected),
		fmt.Sprintf(color.BRed("got: | %v"), got),
		"\n",
		"\n",
	}

	formattedErrorMessage := columnize.SimpleFormat(errorMessage)

	if (expected == got) != expectation {
		t.Errorf(formattedErrorMessage)
		return false
	}
	return true
}

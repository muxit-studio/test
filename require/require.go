package require

import (
	"fmt"

	"github.com/muxit-studio/color"
	"github.com/muxit-studio/columnize"
)

type T interface {
	Fatalf(format string, args ...interface{})
}

func Equal(t T, got interface{}, expected interface{}, message string) {
	require(t, got, expected, message, true)
}

func NotEqual(t T, got interface{}, expected interface{}, message string) {
	require(t, got, expected, message, false)
}

func require(t T, got interface{}, expected interface{}, message string, expectation bool) {
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
		t.Fatalf(formattedErrorMessage)
	}
}

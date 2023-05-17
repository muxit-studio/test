# test

[![Go Report Card](https://goreportcard.com/badge/github.com/muxit-studio/test)](https://goreportcard.com/report/github.com/muxit-studio/test)

`test` is a simple and flexible testing library for Go, designed to work seamlessly with the standard `testing` package provided by Go. It offers additional functionality with color-coded output and columnized messages for easier reading of test results.

## Installation

Use the go get command to add `test` to your project:

```bash
go get github.com/muxit-studio/test
```

## Usage

This package provides two main components - `assert` and `require`.

### assert
```go
import (
	"testing"
	"github.com/muxit-studio/test/assert"
)

func TestSomething(t *testing.T) {
	result := someFunction()
	expected := "expected result"
	assert.Equal(t, result, expected, "Result should match expected value")
}
```

### require
```go
import (
	"testing"
	"github.com/muxit-studio/test/require"
)

func TestSomethingElse(t *testing.T) {
	result := someOtherFunction()
	expected := "expected result"
	require.Equal(t, result, expected, "Result should match expected value")
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

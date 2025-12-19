package currentline

import "runtime"

func CurrentLine() int {
	_, _, line, _ := runtime.Caller(1)
	return line
}

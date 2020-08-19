package log

import (
    "io"
    "fmt"
)
const (
	red    = 31
	yellow = 33
	blue   = 34
	green  = 32
	gray   = 37
)

func printWithColor(w io.Writer, logColor int, msg string) {
	_, _ = fmt.Fprintf(w, "\x1b[1;40;%dm %s \x1b[0m ", logColor, msg)
}

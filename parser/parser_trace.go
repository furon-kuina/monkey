package parser

import (
	"fmt"
	"strings"
)

var nestLevel int

const traceIndentPlaceholder string = "\t"

func indentLevel() string {
	return strings.Repeat(traceIndentPlaceholder, nestLevel-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", indentLevel(), fs)
}

func incIndent() { nestLevel++ }
func decIndent() { nestLevel-- }

func trace(msg string) string {
	incIndent()
	tracePrint("BEGIN " + msg)
	return msg
}

func untrace(msg string) {
	tracePrint("END " + msg)
	decIndent()
}

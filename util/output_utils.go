package util

import (
	"fmt"
)


var headerBounding = "##########"


func PrintFormattedLn(text string, args ...any) {
	if len(args) > 0 {
		fmt.Printf(text, args...)
		fmt.Println()
	} else {
		fmt.Println(text)
	}
}

func FormatHeader(text string) string {
	combined := "\n" + headerBounding + " " + text + " " + headerBounding + "\n"
	return combined
}

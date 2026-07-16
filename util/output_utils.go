package util

import (
	"fmt"
)

var logo = ` _  _   __   __ _  ____  __  
( \/ ) / _\ (  ( \(_  _)/ _\ 
/ \/ \/    \/    /  )( /    \
\_)(_/\_/\_)__) (__)\_/\_/`

var headerBounding = "##########"


func PrintFormattedLn(text string, args ...any) {
	if len(args) > 0 {
		fmt.Printf(text, args...)
		fmt.Println()
	} else {
		fmt.Println(text)
	}
}

// func PrintTitle()

func FormatHeader(text string) string {
	combined := "\n" + headerBounding + " " + text + " " + headerBounding + "\n"
	return combined
}

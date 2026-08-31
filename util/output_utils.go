package util

import (
	"fmt"

	"github.com/fatih/color"
)

const logo = ` _  _   __   __ _  ____  __  
( \/ ) / _\ (  ( \(_  _)/ _\ 
/ \/ \/    \/    /  )( /    \
\_)(_/\_/\_/\_)__) (__)\_/\_/`

const headerBounding = "##########"

const border = "~~~~~~~~~~"


func PrintFormattedLn(text string, args ...any) {
	if len(args) > 0 {
		fmt.Printf(text, args...)
		fmt.Println()
	} else {
		fmt.Println(text)
	}
}

func PrintTitle() {
	PrintFormattedLn(logo)
	PrintFormattedLn("(Created by Henry Allen)")
}

func PrintBorder() {
	PrintFormattedLn("\n" + border + "\n")
}

func FormatHeader(text string) string {
	combined := "\n" + headerBounding + " " + text + " " + headerBounding + "\n"
	// const static = 
	return combined
}

func PrintColor(hex string, text string) {
	r, g, b, err := HexToRGB(hex)
	if err != nil {
		fmt.Println("Error processing group color: ", err)
	}

	colorPrint := color.RGB(r, g, b).PrintfFunc()


	colorPrint(text)
}
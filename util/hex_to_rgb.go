package util

import "fmt"

func HexToRGB(hex string) (int, int, int, error) {
	var r, g, b int

	if len(hex) > 0 && hex[0] == '#' {
		hex = hex[1:]
	}

	if len(hex) != 6 {
		return 0, 0, 0, fmt.Errorf("Invalid hex length.")
	}

	// breaks hex string up into three variables which get read as hexadecimal values with leading zeroes accounted for
	_, err := fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)

	if err != nil {
		return 0, 0, 0, err
	}

	return r, g, b, nil
}
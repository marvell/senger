package senger

import (
	"fmt"
	"os"
)

var (
	// Normal colors
	nBlack   = []byte{'\033', '[', '3', '0', 'm'}
	nRed     = []byte{'\033', '[', '3', '1', 'm'}
	nGreen   = []byte{'\033', '[', '3', '2', 'm'}
	nYellow  = []byte{'\033', '[', '3', '3', 'm'}
	nBlue    = []byte{'\033', '[', '3', '4', 'm'}
	nMagenta = []byte{'\033', '[', '3', '5', 'm'}
	nCyan    = []byte{'\033', '[', '3', '6', 'm'}
	nWhite   = []byte{'\033', '[', '3', '7', 'm'}

	// Bright colors
	bBlack   = []byte{'\033', '[', '3', '0', ';', '1', 'm'}
	bRed     = []byte{'\033', '[', '3', '1', ';', '1', 'm'}
	bGreen   = []byte{'\033', '[', '3', '2', ';', '1', 'm'}
	bYellow  = []byte{'\033', '[', '3', '3', ';', '1', 'm'}
	bBlue    = []byte{'\033', '[', '3', '4', ';', '1', 'm'}
	bMagenta = []byte{'\033', '[', '3', '5', ';', '1', 'm'}
	bCyan    = []byte{'\033', '[', '3', '6', ';', '1', 'm'}
	bWhite   = []byte{'\033', '[', '3', '7', ';', '1', 'm'}

	bWhiteOnRed = []byte{'\033', '[', '3', '7', ';', '1', ';', '4', '1', 'm'}

	reset = []byte{'\033', '[', '0', 'm'}
)

func isTTY() bool {
	if forceColor := os.Getenv("SENGER_FORCE_COLOR"); forceColor != "" {
		if forceColor == "true" || forceColor == "1" {
			return true
		}
	}

	fi, err := os.Stdout.Stat()
	if err == nil {
		m := os.ModeDevice | os.ModeCharDevice
		return fi.Mode()&m == m
	}

	return false
}

func withColor(color []byte, msg string) string {
	if isTTY() {
		return fmt.Sprintf("%s%s%s", color, msg, reset)
	} else {
		return msg
	}
}

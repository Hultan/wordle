package main

import (
	"fmt"
	"os"
	"os/exec"

	faith "github.com/fatih/color"
)

func printWithColor(color colorType, text string) {
	switch color {
	case ColorGreen:
		faith.Set(faith.FgGreen)
	case ColorYellow:
		faith.Set(faith.FgYellow)
	case ColorBlue:
		faith.Set(faith.FgBlue)
	}
	fmt.Print(text)
	faith.Unset()
}

// https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go
func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

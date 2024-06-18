package main

import (
	"fmt"
	"os"
	"strings"

	asciiArt "ascii-art/ascii-funcs-fs"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}

	if len(os.Args) == 3 {
		text := os.Args[1]
		banner := strings.ToLower(os.Args[2])

		// Load ASCII characters from the specified file
		asciiChars, err := asciiArt.LoadAsciiChars(banner + ".txt")
		if err != nil {
			fmt.Println("Error!: Did you mean 'thinkertoy or 'shadow' or 'standard?'", err)
			return
		}
		// Process the first command-line argument with loaded ASCII characters
		asciiArt.ProcessArguments(text, asciiChars)
	}
}

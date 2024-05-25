package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	Args := os.Args
	if len(Args) != 4 {
		fmt.Println("please enter a valid arguments <program Name> <color=COLOR> <letters or words to be colored> <input string>")
		return
	}

	colorArg := os.Args[1]
	color := strings.Split(colorArg, "=")[1] // Extract color from the argument
	coloredLetter := os.Args[2]
	inputedString := os.Args[3]

	for _, v := range inputedString {
		if v < 32 || v > 126 {
			fmt.Println("Please enter a valid input!!")
			return
		}
	}

	count := strings.Count(inputedString, "\\n") // count the number of new lines
	testLines := strings.Split(inputedString, "\\n")
	if len(testLines) == 1 && testLines[0] == "" {
		// Empty input string, print nothing
		return
	}

	content, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	phrase := string(content)
	asciiChars := strings.Split(phrase, "\n\n")
	characters := make([][]string, len(asciiChars))

	for i, char := range asciiChars {
		characters[i] = strings.Split(char, "\n")
	}

	counter := 1
	for _, line := range testLines {
		if line == "" {
			if counter <= count {
				fmt.Println()
			}
			counter++
			continue
		}

		// Process each line
		for l := 0; l < 8; l++ {
			i := 0
			for i < len(line) {
				if line[i] == ' ' {
					fmt.Print(characters[0][l+1]) // 8 spaces for a space character
					i++
					continue
				}else 

				// Check if the current substring matches coloredLetter
				if strings.HasPrefix(line[i:], coloredLetter) {
					// Print colored word
					for j := 0; j < len(coloredLetter); j++ {
						index := line[i+j] - 32
						fmt.Print(colorFunc(color) + characters[index][l] + "\u001b[0m")
					}
					i += len(coloredLetter)
				} else {
					// Print character normally
					index := line[i] - 32
					fmt.Print(characters[index][l])
					i++
				}
			}
			fmt.Println()
		}
	}
}

func colorFunc(color string) string {
	switch color {
	case "white":
		return "\u001b[38;2;255;255;255m"
	case "black":
		return "\u001b[38;2;0;0;0m"
	case "red":
		return "\u001b[38;2;255;0;0m"
	case "green":
		return "\u001b[38;2;0;255;0m"
	case "blue":
		return "\u001b[38;2;0;0;255m"
	case "yellow":
		return "\u001b[38;2;255;255;0m"
	case "pink":
		return "\u001b[38;2;255;0;255m"
	case "grey":
		return "\u001b[38;2;128;128;128m"
	case "purple":
		return "\u001b[38;2;160;32;255m"
	case "brown":
		return "\u001b[38;2;160;128;96m"
	case "orange":
		return "\u001b[38;2;255;160;16m"
	case "cyan":
		return "\u001b[38;2;0;183;235m"
	}
	return "\u001b[38;2;255;255;255m"
}

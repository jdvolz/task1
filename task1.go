package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Hello, World!")
}

func Capitalizenth(s string, nth int) string {
	var b strings.Builder
	var c = 0
	for _, letter := range s {
		if unicode.IsLetter(letter) || unicode.IsNumber(letter) {
			c++
		}
		if c%nth == 0 {
			// capitalize
			b.WriteString(string(unicode.ToUpper(letter)))
		} else {
			// downcase
			b.WriteString(string(unicode.ToLower(letter)))
		}
	}

	return b.String()
}

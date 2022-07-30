package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func Capitalizenth(s string, nth int) string {
	var b bytes.Buffer
	isAlphaNumeric := regexp.MustCompile("^[A-Za-z0-9]+$").MatchString
	var c = 0
	for _, letter := range s {
		if isAlphaNumeric(string(letter)) {
			c++
		}
		if c%nth == 0 {
			// capitalize
			b.WriteString(strings.ToUpper(string(letter)))
		} else {
			// downcase
			b.WriteString(strings.ToLower(string(letter)))
		}
	}

	return b.String()
}

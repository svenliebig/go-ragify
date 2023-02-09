package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

var lower = false

func handleLine(line string) {
	l := ""
	for _, c := range line {
		nextc := c
		if unicode.IsLetter(c) {
			if lower {
				nextc = unicode.ToLower(c)
				lower = false
			} else {
				nextc = unicode.ToUpper(c)
				lower = true
			}
		}

		l = fmt.Sprintf("%s%c", l, nextc)
	}
	fmt.Println(l)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		handleLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

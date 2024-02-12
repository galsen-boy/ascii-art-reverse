package main

import (
	"fmt"
	"os"
)

// pos - symbol in ascii-art file(input.txt). Count - count of lines in ascii-art file(input.txt). Start - number of line in font(standard.txt)
func reverse(font []string, text []string, pos int, count int, start int) {
	if pos != len(text[count]) { //if we are not finish reverse.
		l := len(font[start]) // len of candite for research
		if pos+l <= len(text[count]) {
			if count < 7 {
				if text[count][pos:l+pos] == font[start+count] { // if the font-line and line of ascii-art equal
					reverse(font, text, pos, count+1, start) // compare the next line
				} else {
					reverse(font, text, pos, 0, start+9) // if not equal we try the next symbol in font-file
				}
			} else {
				r := ((start - 1) / 9) + 32 // find and print the letter
				fmt.Printf("%c", r)
				reverse(font, text, pos+l, 0, 1) // we restart our counts for next reserch
			}
		} else {
			reverse(font, text, pos, 0, start+9)
		}
	}
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// the argument must be started from --reverse=
func checkArgs() (bool, string) {
	ok := true
	s := ""
	if len(os.Args) != 2 || len(os.Args[1]) < 9 || os.Args[1][0:10] != "--reverse=" {
		ok = false
		s = "Usage: go run . [OPTION]\n\nEX: go run main.go --reverse=<fileName>\n"
	}
	return ok, s
}

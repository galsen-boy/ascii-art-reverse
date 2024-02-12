package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	ok, er := checkArgs() // the function checks the validity of the arguments
	if ok == true {       // if we get from checkArgs true, work starts =)
		//read font art
		fContent, err := os.ReadFile("standard.txt")
		checkError(err)
		fontData := string(fContent)
		font := strings.Split(fontData, "\n")
		//read art file for reverse
		textFile := os.Args[1][10:]
		textContent, err := os.ReadFile(textFile)
		checkError(err)
		textData := string(textContent)
		text := strings.Split(textData, "\n")
		//If the file contains more than one object ascii-art
		if len(text) > 9 {
			for i := 0; i < len(text)-1; {
				if len(text[i]) > 0 {
					reverse(font, text[i:i+8], 0, 0, 1)
					fmt.Println()
					i = i + 8
					return
				} else {
					fmt.Println()
					i = i + 1
				}
			}
		} else {
			reverse(font, text, 0, 0, 1)
			fmt.Println()
		}
	} else {
		fmt.Print(er)
	}

}

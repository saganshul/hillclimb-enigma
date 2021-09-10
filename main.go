package main

import (
	"bytes"
	"os"
	"regexp"
	"strings"
)

func SanitizePlaintext(plaintext string) string {
	plaintext = strings.TrimSpace(plaintext)
	plaintext = strings.ToUpper(plaintext)
	plaintext = strings.Replace(plaintext, " ", "", -1)
	plaintext = regexp.MustCompile(`[^A-Z]`).ReplaceAllString(plaintext, "X")
	return plaintext
}

func main() {

	if len(os.Args) != 2 {
		panic("Not enough (or too many) command line args")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	var readBuff bytes.Buffer
	_, err = readBuff.ReadFrom(file)
	readTrigramFile()
	if err != nil {
		panic(err)
	}

	cypherText := readBuff.String()
	cypherText = SanitizePlaintext(cypherText)
	HillClimb(cypherText)
	//fmt.Println(cypherText)
	//fmt.Println(CalculateIC(cypherText))
}

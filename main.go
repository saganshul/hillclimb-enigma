// This project repurpose the code of Go Enigma: https://github.com/emedvedev/enigma

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

func PrintConfig(config []RotorConfig, plugBoard Plugboard) string {
	ans := ""
	for _, rotor := range config {
		ans += rotor.ID
		ans += " "
	}
	ans += "\n"
	for _, rotor := range config {
		ans += string(rotor.Start)
		ans += " "
	}
	ans += "\n"

	for i := 0 ; i<26 ; i++ {
		if plugBoard[i] != i && plugBoard[i] > i {
			ans += string(IndexToChar(i))
			ans += string(IndexToChar(plugBoard[i]))
			ans += " "
		}
	}
	ans = ans[:len(ans) - 1]
	return ans
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
	Solve(cypherText)
}

/*func main() {
	Generate()
}*/

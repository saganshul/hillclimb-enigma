package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func SanitizeParagraph(plaintext string) string {
	plaintext = strings.TrimSpace(plaintext)
	plaintext = strings.ToUpper(plaintext)
	plaintext = strings.Replace(plaintext, " ", "", -1)
	plaintext = regexp.MustCompile(`[^A-Z]`).ReplaceAllString(plaintext, "X")
	return plaintext
}

func GenerateRandomConfig() []RotorConfig {
	rand.Seed(time.Now().UTC().UnixNano())
	possibleRotors1 := []string{"Beta", "Gamma"}
	possibleRotors2 := []string{"I", "II", "VI"}

	len1 := len(possibleRotors1)
	len2 := len(possibleRotors2)

	rotors := []string{possibleRotors1[ rand.Intn(len1)], possibleRotors2[rand.Intn(len2)], "IV", "III"}
	rings := []int{1, 1, 1, 16}

	positions := []string{string(rune('A' + rand.Intn(26))), string(rune('A' + rand.Intn(26))), "B", "Q"}

	finalConfig := make([]RotorConfig, 4)
	for index, rotor := range rotors {
		ring := rings[index]
		value := positions[index][0]
		finalConfig[index] = RotorConfig{ID: rotor, Start: value, Ring: ring}
	}
	return finalConfig
}

func GenerateRandomPlugBoard() Plugboard {
	rand.Seed(time.Now().UTC().UnixNano())
	plugBoard := *NewPlugBoard()
	for i := 0 ; i < 10; i++ {
		for ;true; {
			randNum1 := rand.Intn(26)
			randNum2 := rand.Intn(26)
			if plugBoard[randNum1] == randNum1 && plugBoard[randNum2] == randNum2 && randNum1 != randNum2 {
				plugBoard[randNum1] = randNum2
				plugBoard[randNum2] = randNum1
				break
			}
		}
	}
	return plugBoard
}
func Generate() {
	file, err := os.Open("./paragraphs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Error in closing tests file")
		}
	}(file)

	scanner := bufio.NewScanner(file)
	reflector := "C-thin"
	tests := 1
	for scanner.Scan() {
		para := SanitizeParagraph(scanner.Text())
		config := GenerateRandomConfig()
		plugBoard := GenerateRandomPlugBoard()
		e := NewEnigma(config, reflector, plugBoard)

		decoded := e.EncodeString(para)

		f2, _ := os.Create("TESTS/test." + strconv.Itoa(tests) + ".ct.txt")
		f3, _ := os.Create("TESTS/test." + strconv.Itoa(tests) + ".config.txt")

		defer f2.Close()
		defer f3.Close()


		_, err3 := f2.WriteString(decoded)
		_, err4 := f3.WriteString(PrintConfig(config, plugBoard))
		_, err4 = f3.WriteString(para + "\n")

		if err3 != nil {
			log.Fatal(err3)
		}
		if err4 != nil {
			log.Fatal(err4)
		}
		tests++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
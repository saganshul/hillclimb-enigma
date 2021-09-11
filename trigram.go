package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var trigrams = make([]int, 17576)

func CompareTrigramScore(curr float64, new float64) bool {
	return curr < new
}

func GetIntFromTrigram(str string) int {
	return CharToIndex(str[0])*26*26 + CharToIndex(str[1])*26 + CharToIndex(str[2])
}

func readTrigramFile() {
	for i := 0 ; i < 17576 ; i++ {
		trigrams[i]  = 0
	}

	file, err := os.Open("./english_trigrams.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Error in closing trigrams file")
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry := strings.Split(scanner.Text(), " ")
		if len(entry) == 2 {
			intVar, err2 := strconv.Atoi(entry[1])
			if err2 != nil {
				log.Fatal("Error in converting string to int. Line: ", scanner.Text())
			}

			trigrams[GetIntFromTrigram(entry[0])] = intVar
		} else {
			log.Fatal("Error in split line of file. Line: ", scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func CalculateTriGramScore(ct string) float64 {
	var score float64 = 0
	for i:=0 ; i < (len(ct) - 3)  ;  i++ {
		tri := ct[i:i+3]
		score += float64(trigrams[GetIntFromTrigram(tri)])
	}
	return score
}

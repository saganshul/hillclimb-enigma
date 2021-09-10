package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var trigrams map[string]int = make(map[string]int)

func readTrigramFile() {

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
			trigrams[entry[0]] = intVar
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
		if val, ok := trigrams[tri]; ok {
			if ok {
				score += float64(val)
			}
		}
	}
	return score
}

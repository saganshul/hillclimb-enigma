package main

import (
	"strings"
)

func CalculateIC(txt string) float64 {
	var icScore int = 0
	txtLen := len(txt)
	for i := 'A'; i <= 'Z'; i++ {
		freq := strings.Count(txt, string(i))
		icScore += freq*(freq-1)
	}
	return float64(icScore)/float64(txtLen*(txtLen-1))
}

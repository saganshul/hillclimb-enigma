package main

func CalculateIC(txt string) float64 {
	var icScore int = 0
	txtLen := len(txt)
	var freq = make([]int, 26)
	for i:= 0; i < 26; i++ {
		freq[i] = 0
	}
	for i := 0; i < txtLen ; i++ {
		freq[CharToIndex(txt[i])]++
	}
	for i := 0; i < 26; i++ {
		icScore += freq[i]*(freq[i]-1)
	}
	return float64(icScore)/float64(txtLen*(txtLen-1))
}

func CompareICScore(curr float64, new float64) bool {
	if new > 0.07 {
		return false
	} else if curr < new {
		return true
	}
	return false
}

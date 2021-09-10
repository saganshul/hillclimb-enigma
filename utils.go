package main

// CharToIndex returns the alphabet index of a given letter.
func CharToIndex(char byte) int {
	return int(char - 'A')
}

// IndexToChar returns the letter with a given alphabet index.
func IndexToChar(index int) byte {
	return byte('A' + index)
}

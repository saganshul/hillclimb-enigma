package main

func NewPlugBoard() *Plugboard {
	p := Plugboard{}
	for i := 0; i < 26; i++ {
		p[i] = i
	}
	return &p
}

func HillClimbPlugBoard(cypherText string,
	finalConfig []RotorConfig,
	reflector string,
	plugBoard *Plugboard,
	icScore float64,
	f func(string) float64,
	compare func(float64, float64) bool) (float64, *Plugboard) {
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if plugBoard[i] == i && plugBoard[j] == j {
				plugBoard[i] = j
				plugBoard[j] = i
				e := NewEnigma(finalConfig, reflector, *plugBoard)
				decoded := e.EncodeString(cypherText)
				currIcScore := f(decoded)
				if compare(icScore, currIcScore) {
					icScore = currIcScore
				} else {
					plugBoard[i] = i
					plugBoard[j] = j
				}
			} else if plugBoard[i] != i && plugBoard[j] == j {
				tempConnection := plugBoard[i]
				plugBoard[i] = j
				plugBoard[j] = i
				plugBoard[tempConnection] = tempConnection
				e := NewEnigma(finalConfig, reflector, *plugBoard)
				decoded := e.EncodeString(cypherText)
				currIcScore := f(decoded)
				if compare(icScore, currIcScore) {
					icScore = currIcScore
				} else {
					plugBoard[i] = tempConnection
					plugBoard[tempConnection] = i
					plugBoard[j] = j
				}
			} else if plugBoard[i] == i && plugBoard[j] != j {
				tempConnection := plugBoard[j]
				plugBoard[i] = j
				plugBoard[j] = i
				plugBoard[tempConnection] = tempConnection
				e := NewEnigma(finalConfig, reflector, *plugBoard)
				decoded := e.EncodeString(cypherText)
				currIcScore := f(decoded)
				if compare(icScore, currIcScore) {
					icScore = currIcScore
				} else {
					plugBoard[i] = i
					plugBoard[j] = tempConnection
					plugBoard[tempConnection] = j
				}
			} else {
				tempConnection := plugBoard[i]
				tempConnection2 := plugBoard[j]
				plugBoard[i] = j
				plugBoard[j] = i
				plugBoard[tempConnection] = tempConnection
				plugBoard[tempConnection2] = tempConnection2
				e := NewEnigma(finalConfig, reflector, *plugBoard)
				decoded := e.EncodeString(cypherText)
				currIcScore := f(decoded)
				if compare(icScore, currIcScore) {
					icScore = currIcScore
				} else {
					plugBoard[i] = tempConnection
					plugBoard[j] = tempConnection2
					plugBoard[tempConnection] = i
					plugBoard[tempConnection2] = j
				}
			}
		}
	}
	return icScore, plugBoard
}

func HillClimbPlugBoardCombined(cypherText string,
	finalConfig []RotorConfig,
	reflector string) (float64, *Plugboard) {
	finalPlugBoard := NewPlugBoard()
	e := NewEnigma(finalConfig, reflector, *NewPlugBoard())
	decoded := e.EncodeString(cypherText)
	icScore := CalculateIC(decoded)
	icScore, finalPlugBoard = HillClimbPlugBoard(cypherText,
		finalConfig,
		reflector,
		finalPlugBoard,
		icScore,
		CalculateIC,
		CompareICScore)
	// Second Pass with Trigrams
	e = NewEnigma(finalConfig, reflector, *finalPlugBoard)
	triGramScore := CalculateTriGramScore(e.EncodeString(cypherText))
	triGramScore, finalPlugBoard = HillClimbPlugBoard(cypherText,
		finalConfig,
		reflector,
		finalPlugBoard,
		triGramScore,
		CalculateTriGramScore,
		CompareTrigramScore)
	return triGramScore, finalPlugBoard
}

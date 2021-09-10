package main

import (
	"fmt"
)

func NewPlugBoard() *Plugboard {
	p := Plugboard{}
	for i := 0; i < 26; i++ {
		p[i] = i
	}
	return &p
}

func CompareICScore(curr float64, new float64) bool {
	if new > 0.07 {
		return false
	} else if curr < new {
		return true
	}
	return false
}

func CompareTrigramScore(curr float64, new float64) bool {
	return curr < new
}

func HillClimbPlugBoard(cypherText string,
						finalConfig []RotorConfig,
						reflector string,
						plugBoard Plugboard,
						icScore float64,
						f func(string) float64,
						compare func(float64, float64) bool) (float64, *Plugboard) {
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if plugBoard[i] == i && plugBoard[j] == j {
				plugBoard[i] = j
				plugBoard[j] = i
				e := NewEnigma(finalConfig, reflector, plugBoard)
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
				e := NewEnigma(finalConfig, reflector, plugBoard)
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
				e := NewEnigma(finalConfig, reflector, plugBoard)
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
				e := NewEnigma(finalConfig, reflector, plugBoard)
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
	return icScore, &plugBoard
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
		*finalPlugBoard,
		icScore,
		CalculateIC,
		CompareICScore)

	fmt.Println("IcScore: ", icScore)
	fmt.Println("Intermediate Plugboard:")
	for i := 0; i < 26; i++ {
		if  (*finalPlugBoard)[i] != i &&  (*finalPlugBoard)[i] > i {
			i2 :=  (*finalPlugBoard)[i] + 'A'
			fmt.Print(string(i + 'A') + ":"  + string(i2) + ", ")
		}
	}
	fmt.Println()

	// Second Pass with Trigrams
	e = NewEnigma(finalConfig, reflector, *finalPlugBoard)
	triGramScore := CalculateTriGramScore(e.EncodeString(cypherText))
	triGramScore, finalPlugBoard = HillClimbPlugBoard(cypherText,
		finalConfig,
		reflector,
		*finalPlugBoard,
		triGramScore,
		CalculateTriGramScore,
		CompareTrigramScore)
	return triGramScore, finalPlugBoard
}

func HillClimb (cypherText string) {

	possibleRotors1 := []string{"Beta", "Gamma"}
	possibleRotors2 := []string{"I", "II", "VI"}

	rotors := []string{"??", "??", "IV", "III"}
	rings := []int{1, 1, 1, 16}
	positions := []string{"??", "??", "B", "Q"}
	reflector := "C-thin"

	finalConfig := make([]RotorConfig, 4)
	for index, rotor := range rotors {
		ring := rings[index]
		value := positions[index][0]
		finalConfig[index] = RotorConfig{ID: rotor, Start: value, Ring: ring}
	}
	var triGramScore float64
	icScore := CalculateIC(cypherText)
	for _, rotor := range possibleRotors1 {
		for _, rotor2 := range possibleRotors2 {
			for i := 'A'; i <= 'Z'; i++ {
				for j := 'A'; j <= 'Z'; j++ {
					rotors[0] = rotor
					rotors[1] = rotor2
					positions[0] = string(i)
					positions[1] = string(j)
					config := make([]RotorConfig, 4)
					for index, rotor := range rotors {
						ring := rings[index]
						value := positions[index][0]
						config[index] = RotorConfig{ID: rotor, Start: value, Ring: ring}
					}
					e := NewEnigma(config, reflector, *NewPlugBoard())
					decoded := e.EncodeString(cypherText)
					currIcScore := CalculateIC(decoded)
					/*currTriGramScore, currPlugBoard := HillClimbPlugBoardCombined(cypherText, config, reflector)
					if currTriGramScore > triGramScore {
						triGramScore = currTriGramScore
						finalConfig = config
					}*/
					if currIcScore > icScore {
						icScore = currIcScore
						finalConfig = config
					}
				}
			}
		}
	}

	triGramScore, finalPlugBoard := HillClimbPlugBoardCombined(cypherText, finalConfig, reflector)

	fmt.Println("Final Plugboard:")
	for i := 0; i < 26; i++ {
		if  (*finalPlugBoard)[i] != i &&  (*finalPlugBoard)[i] > i {
			i2 :=  (*finalPlugBoard)[i] + 'A'
			fmt.Print(string(i + 'A') + ":"  + string(i2) + ", ")
		}
	}
	fmt.Println()
	fmt.Println(int64(triGramScore))
	fmt.Println(finalConfig)
	e := NewEnigma(finalConfig, reflector, *finalPlugBoard)
	fmt.Println(e.EncodeString(cypherText))


}

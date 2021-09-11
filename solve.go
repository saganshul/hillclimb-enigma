package main

import (
	"fmt"
	"math"
	"sort"
)

type ConfigScore struct {
	config []RotorConfig
	icScore float64
}

func Solve (cypherText string) {
	ctLen := len(cypherText)
	possibleRotors1 := []string{"Beta", "Gamma"}
	possibleRotors2 := []string{"I", "II", "VI"}

	rotors := []string{"??", "??", "IV", "III"}
	rings := []int{1, 1, 1, 16}
	positions := []byte{'?', '?', 'B', 'Q'}
	reflector := "C-thin"

	finalConfig := make([]RotorConfig, 4)
	for index, rotor := range rotors {
		ring := rings[index]
		value := positions[index]
		finalConfig[index] = RotorConfig{ID: rotor, Start: value, Ring: ring}
	}
	finalPlugBoard := NewPlugBoard()
	var triGramScore float64
	var configs[]ConfigScore
	for _, rotor := range possibleRotors1 {
		for _, rotor2 := range possibleRotors2 {
			for i := 0; i < 26; i++ {
				for j := 0; j < 26; j++ {
					rotors[0] = rotor
					rotors[1] = rotor2
					positions[0] = IndexToChar(i)
					positions[1] = IndexToChar(j)
					config := make([]RotorConfig, 4)
					for index, rotor := range rotors {
						ring := rings[index]
						value := positions[index]
						config[index] = RotorConfig{ID: rotor, Start: value, Ring: ring}
					}
					e := NewEnigma(config, reflector, *NewPlugBoard())
					decoded := e.EncodeString(cypherText)
					currIcScore := CalculateIC(decoded)
					configs = append(configs, ConfigScore{config: config, icScore: currIcScore})
				}
			}
		}
	}

	sort.SliceStable(configs, func(i, j int) bool {
		return configs[i].icScore < configs[j].icScore
	})

	// maxTrials I can perform on macbook pro 2019 under 10 minutes
	numTrials := int(math.Max(5, math.Floor(float64(10*60*1000)/ (0.2 * float64(ctLen)))))
	for i := 1 ; i <= numTrials && len(configs) - i >= 0; i++ {
		config := configs[len(configs) - i].config
		currTriGramScore, currPlugBoard := HillClimbPlugBoardCombined(cypherText, config, reflector)
		if currTriGramScore > triGramScore {
			triGramScore = currTriGramScore
			finalConfig = config
			finalPlugBoard = currPlugBoard
		}
	}
	fmt.Println(PrintConfig(finalConfig, *finalPlugBoard))
}

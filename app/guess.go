package app

type Guesser struct {
	Digit1 string
	Digit2 string
	Digit3 string
	Digit4 string
}

func (g Guesser) String() []string {
	return []string{
		g.Digit1,
		g.Digit2,
		g.Digit3,
		g.Digit4,
	}
}

type SecretNumber struct {
	Number       string
	NumberString []string
}

func (sn *SecretNumber) CountCowBull(guess []string) (int, int) {
	var cow, bull int

	if guess[0] == sn.NumberString[0] {
		bull++
	} else if guess[0] == sn.NumberString[1] || guess[0] == sn.NumberString[2] || guess[0] == sn.NumberString[3] {
		cow++
	}

	if guess[1] == sn.NumberString[1] {
		bull++
	} else if guess[1] == sn.NumberString[0] || guess[1] == sn.NumberString[2] || guess[1] == sn.NumberString[3] {
		cow++
	}

	if guess[2] == sn.NumberString[2] {
		bull++
	} else if guess[2] == sn.NumberString[0] || guess[2] == sn.NumberString[1] || guess[2] == sn.NumberString[3] {
		cow++
	}

	if guess[3] == sn.NumberString[3] {
		bull++
	} else if guess[3] == sn.NumberString[0] || guess[3] == sn.NumberString[1] || guess[3] == sn.NumberString[2] {
		cow++
	}

	return cow, bull
}

type UserSecretNumber struct {
	SecretNumber
}

type BotSecretNumber struct {
	SecretNumber
	Cows  map[string]int
	Bulls map[int]string
}

func NewBotSecretNumber() BotSecretNumber {
	return BotSecretNumber{
		Cows:  map[string]int{},
		Bulls: map[int]string{},
	}
}

func (sn *BotSecretNumber) CountCowBull(guess []string) (*Guesser, int, int) {
	var cow, bull int

	if guess[0] == sn.NumberString[0] {
		bull++
		sn.Bulls[0] = guess[0]
		_, ok := sn.Cows[guess[0]]
		if ok {
			delete(sn.Cows, guess[0])
		}
	} else if guess[0] == sn.NumberString[1] || guess[0] == sn.NumberString[2] || guess[0] == sn.NumberString[3] {
		cow++
		sn.Cows[guess[0]] = 0
	}

	if guess[1] == sn.NumberString[1] {
		bull++
		sn.Bulls[1] = guess[1]
		_, ok := sn.Cows[guess[1]]
		if ok {
			delete(sn.Cows, guess[1])
		}
	} else if guess[1] == sn.NumberString[0] || guess[1] == sn.NumberString[2] || guess[1] == sn.NumberString[3] {
		cow++
		sn.Cows[guess[1]] = 1
	}

	if guess[2] == sn.NumberString[2] {
		bull++
		sn.Bulls[2] = guess[2]
		_, ok := sn.Cows[guess[2]]
		if ok {
			delete(sn.Cows, guess[2])
		}
	} else if guess[2] == sn.NumberString[0] || guess[2] == sn.NumberString[1] || guess[2] == sn.NumberString[3] {
		cow++
		sn.Cows[guess[2]] = 2
	}

	if guess[3] == sn.NumberString[3] {
		bull++
		sn.Bulls[3] = guess[3]
		_, ok := sn.Cows[guess[3]]
		if ok {
			delete(sn.Cows, guess[3])
		}
	} else if guess[3] == sn.NumberString[0] || guess[3] == sn.NumberString[1] || guess[3] == sn.NumberString[2] {
		cow++
		sn.Cows[guess[3]] = 3
	}

	if len(sn.Bulls)+len(sn.Cows) == 4 && len(sn.Bulls) < 4 {
		guesser := make(map[int]string, 4)
		for key, val := range sn.Bulls {
			guesser[key] = val
		}

		//reassign value for cow
		tempArr := []int{}
		for _, val := range sn.Cows {
			tempArr = append(tempArr, val)
		}
		tempArr = append(tempArr, tempArr[0])
		tempArr = tempArr[1:]

		for key := range sn.Cows {
			sn.Cows[key] = tempArr[0]
			if len(tempArr) > 1 {
				tempArr = tempArr[1:]
			}
		}

		//build guesser
		for key, val := range sn.Cows {
			index := val
			valOK := false

			for !valOK {
				if guesser[index] != "" {
					if index >= 3 {
						index = 0
					} else {
						index++
					}
				} else {
					guesser[index] = key
					valOK = true
				}
			}
		}

		return &Guesser{
			Digit1: guesser[0],
			Digit2: guesser[1],
			Digit3: guesser[2],
			Digit4: guesser[3],
		}, cow, bull
	}

	return nil, cow, bull
}

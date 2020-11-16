package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/aldoplianto/cow-bull/app"
)

func GenerateBotSecretNumber() string {
	arrSecret := make([]int, 4)
	for i := 0; i <= 3; i++ {
		rand.Seed(time.Now().UnixNano())
		arrSecret[i] = rand.Intn(10)
	}

	return fmt.Sprintf("%d%d%d%d", arrSecret[0], arrSecret[1], arrSecret[2], arrSecret[3])
}

func GenerateGuess() map[int]app.Guesser {
	rand.Seed(time.Now().UnixNano())
	init := rand.Intn(10)
	rand.Seed(time.Now().UnixNano())
	opr := rand.Intn(9)
	if opr%2 == 0 {
		opr++
	}
	if opr == 5 {
		opr = 7
	}

	guessDigit := make([]string, 10)
	if init == 0 {
		init = 1
	}
	for i := range guessDigit {
		if init == 10 {
			guessDigit[i] = strconv.Itoa(0)
			init = opr
		} else {
			guessDigit[i] = strconv.Itoa(init)
			init = init + opr
		}

		if init > 10 {
			init -= 10
		}
	}

	var guessSeq = make(map[int]app.Guesser)
	guessSeq[1] = app.Guesser{
		Digit1: guessDigit[0],
		Digit2: guessDigit[1],
		Digit3: guessDigit[2],
		Digit4: guessDigit[3],
	}

	guessSeq[2] = app.Guesser{
		Digit1: guessDigit[4],
		Digit2: guessDigit[5],
		Digit3: guessDigit[6],
		Digit4: guessDigit[7],
	}

	guessSeq[3] = app.Guesser{
		Digit1: guessDigit[0],
		Digit2: guessDigit[1],
		Digit3: guessDigit[8],
		Digit4: guessDigit[9],
	}

	return guessSeq
}

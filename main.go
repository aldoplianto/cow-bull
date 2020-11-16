package main

import (
	"fmt"

	"github.com/aldoplianto/cow-bull/app"
	gameConst "github.com/aldoplianto/cow-bull/const"
	"github.com/aldoplianto/cow-bull/utils"
)

func main() {

	var guess string
	var guessString []string

	var gameFinished bool
	var inputValidated bool
	var guessValidated bool
	var secretValidated bool

	var guessSeq map[int]app.Guesser
	var guesser *app.Guesser
	var curGuesser app.Guesser
	var cow, bull int
	var secretCow, secretBull int

	var turn int

	fmt.Printf("%s\n\n", gameConst.Intro)

	for {
		userNum := app.UserSecretNumber{}
		botNum := app.NewBotSecretNumber()

		//generate secret number
		for !secretValidated {
			secret := utils.GenerateBotSecretNumber()
			secretString, err := utils.Validate(secret)
			secretValidated = err == nil

			userNum.Number = secret
			userNum.NumberString = secretString
		}

		guessSeq = utils.GenerateGuess()

		//game begin
		for !gameFinished {
			//input user secret number
			for !inputValidated {
				fmt.Printf("%s: ", gameConst.EnterNumber)
				var input string
				fmt.Scanln(&input)
				inputString, err := utils.Validate(input)
				if err != nil {
					fmt.Printf("%v\n\n", err)
				}
				inputValidated = err == nil

				botNum.Number = input
				botNum.NumberString = inputString
			}

			turn++
			fmt.Printf("\n%s: ", gameConst.EnterGuess)
			for !guessValidated {
				var err error
				fmt.Scanln(&guess)
				guessString, err = utils.Validate(guess)
				if err != nil {
					fmt.Printf("%v\n\n", err)
					fmt.Printf("%s: ", gameConst.EnterGuess)
				}
				guessValidated = err == nil
			}

			cow, bull = userNum.CountCowBull(guessString)
			// fmt.Printf("\n%d %d\n\n", cow, bull)

			if guesser != nil {
				curGuesser = *guesser
			} else {
				curGuesser = guessSeq[turn]
			}
			curGuesserString := curGuesser.String()
			guesser, secretCow, secretBull = botNum.CountCowBull(curGuesserString)

			//print turn detail
			fmt.Printf(gameConst.Turn, turn)
			fmt.Printf(gameConst.TurnDetail, guessString, curGuesserString)
			fmt.Printf(gameConst.TurnCowBull, cow, bull, secretCow, secretBull)

			if bull == 4 {
				gameFinished = true
				fmt.Printf(gameConst.Result, userNum.NumberString, botNum.NumberString)
				fmt.Printf(gameConst.Congrats, turn)
			} else {
				guessValidated = false
			}

			if secretBull == 4 {
				gameFinished = true
				fmt.Printf(gameConst.Result, userNum.NumberString, botNum.NumberString)
				fmt.Printf(gameConst.BotCongrats, turn)
			} else {
				guessValidated = false
			}
		}

		fmt.Println(gameConst.TryAgain)
		var tryAgain string
		fmt.Scanln(&tryAgain)
		if tryAgain == "Y" || tryAgain == "y" {
			gameFinished = false
			inputValidated = false
			secretValidated = false
			turn = 0
			continue
		}
		break
	}
}

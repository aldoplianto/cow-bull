package gameConst

const (
	Intro = `--------------- COW BULL ---------------
welcome to cow bull game
you will be playing against bot to guess what number each other has
the number consists of 4 digits, same digits are not allowed
each turn you have to guess the number
if you guess right digit on the wrong place, you will get cow
if you guess right digit on the right place, you will get bull
whoever is the fastest to guess the opponent's number will win the game
please enjoy`

	EnterNumber = "please enter your secret number"
	EnterGuess  = "please enter your guess"

	Congrats    = "Congratulation\nyou have completed this game in %d turn\n"
	BotCongrats = "Aw sorry, you lost to bot\nbot has completed this game in %d turn\n"
	Result      = "\nbot secret number: %s\tyour secret number: %s\n\n"

	Turn        = "\n--------------- TURN %d ---------------\n"
	TurnDetail  = "your guess\t\t\tbot guess\n%s\t\t\t%s\n"
	TurnCowBull = "Cow: %d, Bull: %d\t\t\tCow: %d, Bull: %d\n\n"

	TryAgain = "\nDo you want to try again? (Y/N)"
)

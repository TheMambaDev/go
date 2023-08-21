package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	number, err := GuessTheNumber()
	if err != nil {
		fmt.Print(err)
	}

	var guess int64

	fmt.Print("What number did you guess (0 - 10): ")
	fmt.Scan(&guess)

	isCorrect := func() bool {
		return guess == number
	}

	if isCorrect() {
		fmt.Printf("You're right the number is: %v \n", guess)
	} else {

		var tryAgain int
		fmt.Printf("unfortunately, the number was %v. want to try again ? (0 or 1)", number)

		fmt.Scan(&tryAgain)

		switch tryAgain {
		case 0:
			return
		case 1:
			main()

		default:
			return
		}

	}
}

func GuessTheNumber() (number int64, err error) {
	num, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		return 0, err
	}

	number = num.Int64()
	return number, err
}

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	Easy = iota + 1
	Medium
	Hard
)

type GameConfig struct {
	lives         int
	description   string
	difficultyStr string
}

func getDifficultyConfig(level int) GameConfig {
	configs := map[int]GameConfig{
		Easy:   {lives: 10, description: "Easy", difficultyStr: "You have 10 chances to guess the correct number."},
		Medium: {lives: 5, description: "Medium", difficultyStr: "You have 5 chances to guess the correct number."},
		Hard:   {lives: 3, description: "Hard", difficultyStr: "You have 3 chances to guess the correct number."},
	}

	if config, ok := configs[level]; ok {
		return config
	}
	return GameConfig{}
}

func readInt(prompt string) (int, error) {
	var num int
	fmt.Print(prompt)
	_, err := fmt.Scanf("%d", &num)
	return num, err
}

func getDifficultyChoice() int {
	fmt.Println("\nWelcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("\nPlease select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")

	for {
		choice, err := readInt("\nEnter your choice (1-3): ")
		if err == nil && choice >= Easy && choice <= Hard {
			return choice
		}
		fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
	}
}

func playGame(targetNum, lives int) bool {
	for i := 1; i <= lives; i++ {
		remainingLives := lives - i + 1
		fmt.Printf("\nAttempt %d/%d\n", i, lives)
		
		guess, err := readInt("Enter your guess: ")
		if err != nil {
			fmt.Println("Please enter a valid number between 1 and 100.")
			i-- 
			continue
		}

		if guess < 1 || guess > 100 {
			fmt.Println("Please enter a number between 1 and 100.")
			i-- 
			continue
		}

		if guess == targetNum {
			fmt.Printf("\nCongratulations! You guessed the correct number in %d attempts!\n", i)
			return true
		}

		hint := "greater"
		if guess > targetNum {
			hint = "less"
		}
		fmt.Printf("Incorrect! The number is %s than %d.\n", hint, guess)
		
		if remainingLives > 1 {
			fmt.Printf("You have %d attempts remaining.\n", remainingLives-1)
		}
	}
	return false
}

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		choice := getDifficultyChoice()
		config := getDifficultyConfig(choice)

		fmt.Printf("\nGreat! You have selected the %s difficulty level.\n", config.description)
		fmt.Println(config.difficultyStr)
		fmt.Println("Let's start the game!")

		targetNum := rng.Intn(100) + 1
		won := playGame(targetNum, config.lives)

		if !won {
			fmt.Printf("\nGame Over! The number was %d.\n", targetNum)
		}

		fmt.Print("\nWould you like to play again? (y/n): ")
		var playAgain string
		fmt.Scanf("%s", &playAgain)
		if playAgain != "y" && playAgain != "Y" {
			fmt.Println("Thanks for playing! Goodbye!")
			os.Exit(0)
		}
	}
}
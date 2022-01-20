package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	faith "github.com/fatih/color"
)

func main() {
	words := getWords()

	// for {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(words))
	word := words[r]
	for {
		fmt.Print("Guess : ")
		guess, _ := reader.ReadString('\n')
		guess = guess[:5]
		evaluateGuess(word, guess)
		if guess == word {
			fmt.Println("You won!")
			break
		}
	}
	fmt.Println("The result word was : ", word)
	// }
}

func evaluateGuess(word string, guess string) {
	for i := range guess {
		if guess[i] == word[i] {
			faith.Set(faith.FgGreen)
			fmt.Print(string(guess[i]))
			faith.Unset()
		} else if strings.Contains(word, string(guess[i])) {
			faith.Set(faith.FgYellow)
			fmt.Print(string(guess[i]))
			faith.Unset()
		} else {
			fmt.Print(string(guess[i]))
		}
	}

	fmt.Println()
}

func getWords() []string {
	var words []string

	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return words
}

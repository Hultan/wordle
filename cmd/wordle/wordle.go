package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	faith "github.com/fatih/color"
)

type colorType int

const (
	ColorNone colorType = iota
	ColorGreen
	ColorYellow
	ColorBlue
)

var words []string
var guesses []string

func main() {
	words = getWords()

	// for {
	clear()
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(words))
	word := words[r]
	for {
		fmt.Print("Next guess : ")
		guess, _ := reader.ReadString('\n')
		guess = guess[:len(guess)-1] // Ignore the new line character

		if guess == "?" {
			break // Give up
		}
		if len(guess) < 5 {
			fmt.Println("Need 5 letters...")
			continue
		} else {
			guess = guess[:5] // Use the first 5 letters as a guess
		}
		if !contains(words, guess) {
			fmt.Printf("'%s' a word, try again...\n", guess)
			continue
		}
		evaluateGuess(word, guess)
		if guess == word {
			fmt.Printf("You won in %d guesses!\n", len(guesses))
			break
		}
	}
	fmt.Print("The word was : ")
	faith.Set(faith.FgBlue)
	fmt.Println(word)
	faith.Unset()
	// }
}

func evaluateGuess(word string, guess string) {
	clear()

	guesses = append(guesses, guess)
	for i := range guesses {
		printGuess(i, guesses[i], word)
	}
}

func printGuess(num int, guess, word string) {
	printWithColor(ColorBlue, fmt.Sprintf("%s\t: ", strconv.Itoa(num+1)))

	for i := range guess {
		s := string(guess[i])

		if guess[i] == word[i] {
			printWithColor(ColorGreen, s)
		} else if strings.Contains(word, s) {
			printWithColor(ColorYellow, s)
		} else {
			printWithColor(ColorNone, s)
		}
	}

	fmt.Println()
}

func getWords() []string {
	var list []string

	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	return list
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"

	faith "github.com/fatih/color"
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
		guess = guess[:5]
		if !contains(words, guess) {
			fmt.Println("Not a word, try again...")
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
	faith.Set(faith.FgBlue)
	fmt.Printf("%s\t: ", strconv.Itoa(num+1))
	faith.Unset()

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

// https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go
func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
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
		return list[i]<list[j]
	})

	return list
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
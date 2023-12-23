package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	var words []string
	var correctAnswers int
	var timeStart = time.Now()

	var file, err = os.Open("words.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	for time.Now().Before(timeStart.Add(time.Minute)) {
		remainingTime := timeStart.Add(time.Minute).Sub(time.Now())

		var word = words[rand.Intn(len(words))]

		fmt.Printf("You have %.0f seconds left\n", remainingTime.Seconds())
		fmt.Println(word)

		var answer string
		fmt.Print("Your word: ")
		fmt.Scanln(&answer)

		if strings.TrimSpace(word) == strings.TrimSpace(answer) {
			correctAnswers++
		}
	}

	var wpm = float64(correctAnswers) / 60 * 100
	fmt.Printf("\nWords per minute: %d\n", int(math.Ceil(wpm)))
}

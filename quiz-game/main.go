package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 [ ] TODO : Add Timer
 [ ] TODO : Add Flags
*/

func readCSV(fileName string) [][]string {
	// Doc : https://golangdocs.com/reading-and-writing-csv-files-in-golang

	fmt.Printf("reading %s\n", fileName)
	rawQuizContent, quizFileErr := os.Open(fileName)
	if quizFileErr != nil {
		fmt.Printf("Error reading quiz file : %v", quizFileErr)
	}
	reader := csv.NewReader(rawQuizContent)
	records, _ := reader.ReadAll()

	fmt.Printf("Testing %v\n", records)
	fmt.Println("Questions loaded")
	return records
}

// promptQuestions takes in the slice of questions and prompts the user
// returns an integer of the number of correct questions
func promptQuestions(questions [][]string) (int, int) {
	numCorrect := 0
	totalQuestions := 0
	answerPrompt := bufio.NewReader(os.Stdin)
	// Reading from stdin
	// https://stackoverflow.com/questions/20895552/how-to-read-from-standard-input-in-the-console
	for i, record := range questions {
		fmt.Printf("Question #%v\n  %v:\n", i+1, record[0])
		answer, _ := answerPrompt.ReadString('\n')
		// Needed to remove the newline from the input
		// https://golang.org/pkg/strings/#TrimRight
		responseInt, responseErr := strconv.Atoi(strings.TrimRight(answer, "\n"))
		if responseErr != nil {
			// if unable to convert user input to int, consider answer incorrect
			fmt.Printf("Unable to convert %v to int\n", answer)
			totalQuestions++
			continue
		}
		answerInt, answerErr := strconv.Atoi(record[1])
		if answerErr != nil {
			// if unable to convert str to int, answer is considered invalid and skipped
			fmt.Printf("Unable to convert %v to int, skipping question\n", record[1])
			continue
		}
		if responseInt == answerInt {
			fmt.Println("Correct!")
			numCorrect++
			totalQuestions++
		} else {
			fmt.Println("Incorrect.")
			totalQuestions++
		}
	}
	return numCorrect, totalQuestions
}

func main() {
	quizFile := "quiz.csv"
	fmt.Println("Welcome to the Math Quiz")
	quizContent := readCSV(quizFile)
	// for i, record := range quizContent {
	// 	fmt.Printf("Record : %v\n", record)
	// 	fmt.Printf("Question %v : %v\n", i+1, record[0])
	// 	fmt.Printf("Answer : %v\n", record[1])
	// }
	correctAnswers, totalQuestions := promptQuestions(quizContent)
	fmt.Printf("Answered %v out of %v\n", correctAnswers, totalQuestions)
}

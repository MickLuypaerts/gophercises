package main

import (
	"os"
	"encoding/csv"
	"fmt"
	"log"
	"strconv"
	"time"
)

type QuestionAnswer struct {
	question string
	answer string
}

type UserScore struct {
	correct int
	incorrect int
}

func (u UserScore) Total() int {
	return u.correct + u.incorrect
}

func main() {
	records, err := readCsv("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	questionAnswerSlice := convertRecordsToQAStruct(records)
	var userScore UserScore
	startTimer(5, &userScore)
	userScore = askQuestionsGetAnswers(questionAnswerSlice)
	fmt.Printf("Correct Answers: %d\nTotal number of questions: %d\n", userScore.correct, userScore.Total())
}

func startTimer(timeForQuiz int, userScore *UserScore) {
	fmt.Printf("Press enter to start the quiz you will have %d seconds to complete it", timeForQuiz)
	fmt.Scanln()
	go timer(timeForQuiz, userScore)
}

func timer(timeForQuiz int, userScore *UserScore) {
	timer := time.NewTimer(time.Duration(timeForQuiz) * time.Second)
	<-timer.C
	fmt.Printf("Time ran out.\nCorrect Answers: %d\nTotal number of questions: %d\n", userScore.correct, userScore.Total())
	os.Exit(0)
}

func readCsv(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return records, nil
}

func convertRecordsToQAStruct(records [][]string) ([]QuestionAnswer) {
	questionAnswerSlice := make([]QuestionAnswer, len(records))
	for i, record := range records {
		questionAnswer := QuestionAnswer{
			question: record[0],
			answer: record[1],
		}
		questionAnswerSlice[i] = questionAnswer
	}
	return questionAnswerSlice
}

func askQuestionsGetAnswers(questionAnswerSlice []QuestionAnswer) UserScore {
	var userScore UserScore
	for _, questionAnswer := range questionAnswerSlice {
		var userInput int
		fmt.Printf("%s = ", questionAnswer.question)
		fmt.Scan(&userInput)

		if answer, _ := strconv.Atoi(questionAnswer.answer); answer == userInput {
			userScore.correct++
		} else {
			userScore.incorrect++
		}
	}
	return userScore
}

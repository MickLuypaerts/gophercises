package main

import (
	"os"
	"encoding/csv"
	"fmt"
	"log"
	"strconv"
	"time"
	"flag"
	"math/rand"
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

	timeFlag := flag.Int("time", 30, "Time to complete the quiz.")
	shuffleFlag := flag.Bool("shuffle", false, "Should the quiz be shuffeld.")
	flag.Parse()

	records, err := readCsv("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	questionAnswerSlice := convertRecordsToQAStruct(records)

	if *shuffleFlag {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(questionAnswerSlice), func(i, j int) {questionAnswerSlice[i], questionAnswerSlice[j] = questionAnswerSlice[j], questionAnswerSlice[i] })
	}

	var userScore UserScore
	startTimer(*timeFlag, &userScore)
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

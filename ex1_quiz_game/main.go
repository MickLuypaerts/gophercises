package main

import (
	"os"
	"encoding/csv"
	"fmt"
	"log"
)

type QuestionAnswer struct {
	question string
	answer string
}
func main(){
	partOne()
}

func partOne(){
	records, err := readCsv("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	questionAnswerSlice := convertRecordsToQAStruct(records)
	fmt.Println(questionAnswerSlice)
}

func readCsv(fileName string) ([][]string, error){
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
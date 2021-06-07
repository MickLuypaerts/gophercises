package main

import (
	"os"
	"encoding/csv"
	"fmt"
)

type QuestionAnswer struct {
	question string
	answer string
}
func main(){
	partOne()
}

func partOne(){
	readCsv("problems.csv")
}

func readCsv(fileName string) ([][]string, error){
	file, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	record, err := reader.Read()
	if err != nil {
		return [][]string{}, err
	}
	fmt.Println(record)
	record, err = reader.Read()
	fmt.Println(record[1])

	return nil, err
}
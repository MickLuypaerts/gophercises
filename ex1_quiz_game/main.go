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
	records, err := reader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	fmt.Println(records)

	return records, nil
}
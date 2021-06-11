package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option
}

func main() {

	parsedJSON, err := parseJSON("gopher.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(parsedJSON)

}

func parseJSON(file string) (map[string]StoryArc, error) {

	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	storyMap := make(map[string]StoryArc)
	json.Unmarshal([]byte(byteValue), &storyMap)

	return storyMap, nil
}

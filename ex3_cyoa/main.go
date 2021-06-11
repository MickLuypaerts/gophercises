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
	Options []Option //`json:"options"`
}

func main() {

	parsedJSON, err := parseJSON("gopher.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(parsedJSON["intro"].Options)

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
	/*
		Unmarshal operates on a []byte, meaning that it needs the JSON to be fully loaded in memory. If
		you already have the full JSON stored in a variable this will likely be a bit faster.
		Decoder operates over a stream, any object that implements the io.Reader interface. Meaning
		that you can parse the JSON as its being received/transmitted without having to fully store it in
		memory. This is useful when dealing with a large dataset by not requiring you to create an extra
		copy of the entire JSON content in memory.
	*/
	if err := json.Unmarshal([]byte(byteValue), &storyMap); err != nil {
		return nil, err
	}

	return storyMap, nil
}

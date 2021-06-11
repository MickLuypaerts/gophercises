package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var storyTmpl = `
<!DOCTYPE html>
<html>
<head>
    <title>Choose your own adventure</title>
</head>
<style>
	h1 {
		background-color: coral;
	}
</style>
<body>
    <h1>{{.Title}}</h1>
    {{range .Story}}
    <p>{{.}}</p>
    {{end}}

    <ul>
        {{range .Options}}
        <li><a href="/{{.Arc}}">{{.Text}}</a></li>
        {{end}}
    </ul>
</body>
</html>`

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type handler struct {
	story map[string]StoryArc
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	tmp := template.Must(template.New("").Parse(storyTmpl))

	path := r.URL.Path

	if path == "" || path == "/" {
		err := tmp.Execute(w, h.story["intro"])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(strings.Trim(path, "/"))
		err := tmp.Execute(w, h.story[strings.Trim(path, "/")])
		if err != nil {
			log.Fatal(err)
		}
	}

}

func main() {

	parsedJSON, err := parseJSON("gopher.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Starting the server on :8080")
	log.Fatal(http.ListenAndServe(":8080", MyHandler(parsedJSON)))

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

func MyHandler(s map[string]StoryArc) http.Handler {
	return handler{s}
}

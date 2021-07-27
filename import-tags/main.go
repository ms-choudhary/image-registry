package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var allTags map[string]string

func loadTags() {
	filepath := os.Getenv("TAGS_FILE_PATH")
	if filepath == "" {
		fmt.Fprintf(os.Stderr, "Please set TAGS_FILE_PATH\n")
		os.Exit(1)
	}

	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %s: %v\n", filepath, err)
		os.Exit(1)
	}

	err = json.Unmarshal(bytes, &allTags)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse json: %v\n", err)
		os.Exit(1)
	}
}

func prevTag(key string) string {
	val := allTags[key]
	if val == "" {
		return "latest"
	}
	return val
}

func main() {
	loadTags()
	fmt.Println(prevTag("default/golang"))
	fmt.Println(prevTag("default/golang-ex"))
	fmt.Println(prevTag("default/busybox"))
}

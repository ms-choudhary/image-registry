package imagestream

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var allTags map[string]string

func LoadTags() {
	filepath := os.Getenv("TAGS_FILE_PATH")
	if filepath == "" {
		fmt.Fprintf(os.Stderr, "Please set TAGS_FILE_PATH\n")
		os.Exit(1)
	}

	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[tags] failed to open file: %s: %v\n", filepath, err)
		os.Exit(1)
	}

	err = json.Unmarshal(bytes, &allTags)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[tags] failed to parse json: %v\n", err)
		os.Exit(1)
	}
}

func PrevTag(key string) string {
	val := allTags[key]
	if val == "" {
		fmt.Fprintf(os.Stderr, "[tags] key missing or empty: %s\n", key)
	}
	return val
}

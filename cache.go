package main

import (
	"bufio"
	"os"
	"slices"
	"strings"
)

func GetCache(filePath string) (tweets []Tweet) {
	file, err := os.Open(filePath)

	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		tweets = append(tweets, Tweet{Url: scanner.Text()})
	}

	return

}

func WriteCache(filePath string, cache []Tweet) (err error) {

	textList := slices.Collect(func(yeld func(string) bool) {
		for _, tweet := range cache {
			if !yeld(tweet.Url) {
				return
			}
		}
	})

	text := strings.Join(textList, "\n")

	err = os.WriteFile(
		filePath,
		[]byte(text),
		os.ModePerm|os.ModeAppend,
	)

	return

}

package cache

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strings"
	"sync"

	"github.com/CyberDruga/twitter-bot/src/models"
)

var Tweets = &[]models.Tweet{}

func AddTweet(tweet models.Tweet) {
	*Tweets = append(*Tweets, tweet)
}

var mutex = &sync.Mutex{}

func Lock() {
	slog.Debug("Locking cache")
	mutex.Lock()
}

func Unlock() {
	slog.Debug("Unlocking cache")
	mutex.Unlock()
}

func LoadCache(filePath string) (err error) {

	slog.Debug("Loading cache")
	defer slog.Debug(fmt.Sprintf("Done loading cache: Error %v", err))

	file, err := os.Open(filePath)

	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		*Tweets = append(*Tweets, models.Tweet{Url: scanner.Text()})
	}

	slog.Debug(fmt.Sprintf("cache: %v", *Tweets))

	slog.Debug(fmt.Sprintf("Tweets: %d", len(*Tweets)))

	return

}

func SaveCache(filePath string) (err error) {

	slog.Debug("Saving cache")
	defer slog.Debug(fmt.Sprintf("Done saving cache. Error: %v", err))

	textList := slices.Collect(func(yeld func(string) bool) {
		for _, tweet := range *Tweets {
			if !yeld(tweet.Url) {
				return
			}
		}
	})

	text := strings.Join(textList, "\n")

	err = os.WriteFile(
		filePath,
		[]byte(text),
		os.ModePerm,
	)

	return

}

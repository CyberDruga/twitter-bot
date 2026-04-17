package cache

import (
	"bufio"
	"github.com/charmbracelet/log"
	"os"
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
	log.Debug("Locking cache")
	mutex.Lock()
}

func Unlock() {
	log.Debug("Unlocking cache")
	mutex.Unlock()
}

func LoadCache(filePath string) (err error) {

	log.Debug("Loading cache")
	defer log.Debug("Done loading cache", "Error", err)

	file, err := os.Open(filePath)

	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		*Tweets = append(*Tweets, models.Tweet{Url: scanner.Text()})
	}

	log.Debug(
		"",
		"Tweets", len(*Tweets),
		"Cache", *Tweets,
	)

	return

}

func SaveCache(filePath string) (err error) {

	log.Debug("Saving cache")
	defer log.Debug("Done saving cache", "Error", err)

	var textList []string

	for _, tweet := range *Tweets {
		textList = append(textList, tweet.Url)
	}

	text := strings.Join(textList, "\n")

	err = os.WriteFile(
		filePath,
		[]byte(text),
		os.ModePerm,
	)

	return

}

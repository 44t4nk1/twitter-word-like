package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/44t4nk1/twitter-word-like/api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func GetTweets(c *fiber.Ctx) error {
	username := c.Params("user", "")
	if username == "" {
		return c.Status(404).SendString("Not Found")
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	data := GetTwitterID(os.Getenv("TWITTER_URL") + "by/username/" + username)

	tweets := GetTweetsByID(os.Getenv("TWITTER_URL") + data.Data.ID + "/tweets?exclude=retweets,replies&tweet.fields=public_metrics&max_results=100")

	sanitisedTweets := SanitiseText(tweets)

	tweetWordLike := make(map[string]int)

	for _, tweet := range sanitisedTweets {
		individualTweetText := strings.Split(tweet.Text, " ")
		for _, word := range individualTweetText {
			if val, ok := tweetWordLike[word]; ok {
				val += tweet.LikeCount
				tweetWordLike[word] = val
			} else {
				tweetWordLike[word] = tweet.LikeCount
			}
		}
	}

	p := make(models.PairList, len(tweetWordLike))

	i := 0
	for k, v := range tweetWordLike {
		p[i] = models.Pair{Key: k, Value: v}
		i++
	}

	sort.Sort(p)

	return c.Status(200).JSON(p)
}

func SanitiseText(tweets []models.UserTweet) []models.UserTweetData {

	var cleanTweets []models.UserTweetData
	for _, tweet := range tweets {
		var tweetData models.UserTweetData
		tweetData.LikeCount = tweet.PublicMetrics.LikeCount
		cleanTweet := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(tweet.Text, "\n", ""), "“", ""), ",", ""), ".", ""), "’", ""), "?", ""), "!", ""), "”", ""), "&gt;", ""))
		tweetData.Text = cleanTweet
		tweetData.LikeCount = tweet.PublicMetrics.LikeCount
		cleanTweets = append(cleanTweets, tweetData)
	}

	return cleanTweets
}

func GetTweetsByID(url string) []models.UserTweet {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	req.Header.Set("Authorization", os.Getenv("BEARER"))
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	var data models.UserTweetList

	json.NewDecoder(res.Body).Decode(&data)

	nextToken := data.Meta.NextToken

	var allTweets []models.UserTweet
	allTweets = append(allTweets, data.Data...)

	for nextToken != "" {
		val := GetMoreTweets(url + "&pagination_token=" + nextToken)
		allTweets = append(allTweets, val.Data...)
		nextToken = val.Meta.NextToken
	}

	if err != nil {
		log.Fatalln(err)
	}

	return allTweets
}

func GetMoreTweets(url string) models.UserTweetList {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	req.Header.Set("Authorization", os.Getenv("BEARER"))
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	var data models.UserTweetList

	json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		log.Fatalln(err)
	}

	return data
}

func GetTwitterID(url string) models.UserTwitterBase {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	req.Header.Set("Authorization", os.Getenv("BEARER"))
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	var data models.UserTwitterBase

	json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		log.Fatalln(err)
	}

	return data
}

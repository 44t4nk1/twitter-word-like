package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

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

	url := os.Getenv("TWITTER_USERNAME_URL") + username

	data := GetTwitterID(url)

	return c.Status(200).JSON(data.Data)
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

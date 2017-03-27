package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	incomingWebhookURL  string = "YOUR_WEBHOOK_API_HERE"
	githubRepositoryURL string = "https://api.github.com/repositories"
)

func fetchPublicRepositories() ([]Repository, error) {
	// Github API requires a user-agent :D
	req, err := http.NewRequest("GET", githubRepositoryURL, nil)
	req.Header.Set("User-Agent", "request")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var data []Repository
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data
}

func postToSlack(channel, text, iconEmoji, username string) (bool, error) {
	jsonStr, err := json.Marshal(Message{
		Channel:   channel,
		Text:      text,
		IconEmoji: iconEmoji,
		Username:  username,
	})

	req, err := http.NewRequest("POST", incomingWebhookURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return true, nil
}

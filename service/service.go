package service

import (
	"bytes"
	"encoding/json"
	// "fmt"
	"github.com/alextanhongpin/notification-challenge/common"
	"github.com/alextanhongpin/notification-challenge/model"
	"io/ioutil"
	"net/http"
)

const (
	githubRepositoryURL string = "https://api.github.com/repositories"
)

var configuration = common.GetConfig()

func FetchPublicRepositories() ([]model.Repository, error) {
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

	var data []model.Repository
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func PostToSlack(slack_channel, text, iconEmoji, username string) (bool, error) {
	jsonStr, err := json.Marshal(model.Message{
		Channel:   slack_channel,
		Text:      text,
		IconEmoji: iconEmoji,
		Username:  username,
	})

	req, err := http.NewRequest("POST", configuration.SlackWebhookURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		return false, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)

	return true, nil
}

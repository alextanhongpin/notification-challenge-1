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

// FetchPublicRepositories fetch a list of public repository
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

	// Write to JSON file
	// jsonOutput, _ := json.Marshal(data)
	// err = ioutil.WriteFile("github_repos.json", jsonOutput, 0644)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// PostToSlack post a message to Slack's webhook
func PostToSlack(request interface{}) (bool, error) {
	payload := request.(model.Message)
	jsonStr, err := json.Marshal(payload)

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

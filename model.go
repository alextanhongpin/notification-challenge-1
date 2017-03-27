package main

// Message is the payload for Slack API
type Message struct {
	Channel   string `json:"channel"`
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
}

// Repository is the schema for the Github Repository
type Repository struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	HTMLURL     string `json:"html_url"`
	Description string `json:"description"`
}

// Configuration for the application
type Configuration struct {
	SlackWebhookURL string `json:"slack_webhook_url"`
}

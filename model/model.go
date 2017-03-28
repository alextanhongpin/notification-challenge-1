package model

import (
	"time"
)

// Attachment belongs to Message
type Attachment struct {
	Fallback   string `json:"fallback"`
	Color      string `json:"color, omitempty"`
	Pretext    string `json:"pretext,omitempty"`
	AuthorName string `json:"author_name"`
	AuthorLink string `json:"author_link"`
	AuthorIcon string `json:"author_icon"`
	Title      string `json:"title"`
	TitleLink  string `json:"title_link"`
	Text       string `json:"text"`
	ImageURL   string `json:"image_url"`
	ThumbURL   string `json:"thumb_url"`
	Footer     string `json:"footer"`
	FooterIcon string `json:"footer_icon"`
	Timestamp  int64  `json:"ts"`
}

// Message is the payload for Slack API
type Message struct {
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	Username    string       `json:"username"`
	IconEmoji   string       `json:"icon_emoji"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

// Owner is the owner object
type Owner struct {
	AvatarURL string `json:"avatar_url"`
	URL       string `json:"url"`
	Login     string `json:"login"`
}

// FetchPublicRepositoriesResponse is the response from the service
type FetchPublicRepositoriesResponse struct {
	TotalCount        int          `json:"total_count"`
	IncompleteResults bool         `json:"incomplete_results"`
	Items             []Repository `json:"items"`
}

// Repository is the schema for the Github Repository
type Repository struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	FullName    string    `json:"full_name"`
	HTMLURL     string    `json:"html_url"`
	Description string    `json:"description"`
	Owner       Owner     `json:"owner"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Configuration for the application
type Configuration struct {
	SlackWebhookURL string `json:"slack_webhook_url"`
}

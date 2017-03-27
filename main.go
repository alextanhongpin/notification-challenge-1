package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/robfig/cron.v2"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
)

// Message is the payload for Slack API
type Message struct {
	Channel   string `json:"channel"`
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
}

const incomingWebhookURL string = "YOUR_WEBHOOK_API_HERE"

func fetch() {
	jsonStr, err := json.Marshal(Message{
		Channel:   "#general",
		Text:      "Hello World",
		IconEmoji: ":taco:",
		Username:  "alexbot",
	})
	fmt.Println(jsonStr)
	// var jsonStr = []byte(`
	//    {
	//      "text": "Hello World",
	//      "channel": "#general",
	//      "icon_emoji": ":ghost:",
	//      "username": "alextanhongpin"
	//    }
	//  `)
	req, err := http.NewRequest("POST", incomingWebhookURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("User-Agent", "request")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))
}

// const oneSecond = 1*time.Second + 10*time.Millisecond

func main() {
	fmt.Println("Start Task!")
	c := cron.New()
	c.AddFunc("0 */5 * * * *", func() {
		fmt.Println("Hello world!")
		fetch()
	})
	c.Start()
	fmt.Println("End Task!")
	defer c.Stop()

	// Let the process runs forever
	a := make(chan os.Signal)
	signal.Notify(a, os.Interrupt, os.Kill)
	<-a
}

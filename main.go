package main

import (
	"fmt"
	"gopkg.in/robfig/cron.v2"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("Started cron job")
	fmt.Println(greet())
	config := getConfig()
	fmt.Printf("Printing the URL:%s", config.SlackWebhookURL)
	c := cron.New()
	c.AddFunc("0 */5 * * * *", func() {
		fmt.Println("Hello world!")
		// postToSlack()
	})
	c.Start()
	defer c.Stop()

	// fetchPublicRepositories()

	// Let the process runs forever
	a := make(chan os.Signal)
	signal.Notify(a, os.Interrupt, os.Kill)
	<-a
}

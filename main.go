package main

import (
	"fmt"
	"gopkg.in/robfig/cron.v2"

	"os"
	"os/signal"
)

func main() {
	fmt.Println("Started cron job")

	c := cron.New()
	c.AddFunc("0 */5 * * * *", func() {
		fmt.Println("Hello world!")
		postToSlack()
	})
	c.Start()
	defer c.Stop()

	fetchPublicRepositories()

	// Let the process runs forever
	a := make(chan os.Signal)
	signal.Notify(a, os.Interrupt, os.Kill)
	<-a
}

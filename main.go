package main

import (
	"fmt"
	"github.com/alextanhongpin/notification-challenge/common"
	"github.com/alextanhongpin/notification-challenge/service"
	"gopkg.in/robfig/cron.v2"
	"os"
	"os/signal"
)

var cache map[int]bool

func main() {
	fmt.Println("Started cron job")

	if cache == nil {
		cache = make(map[int]bool)
	}

	c := cron.New()
	c.AddFunc("0 */5 * * * *", func() {
		fmt.Println("Hello world!")
		// service.PostToSlack("#general", "Greetings from go :grin:", ":ghost:", "alextanhongpin")

		repos, err := service.FetchPublicRepositories()
		if err != nil {
			fmt.Println(err)
		}
		slicedRepos := repos[0:5]
		notificationPayload := common.MakeNotificationPayload(cache, slicedRepos)
		similar := common.GetSimilarData(cache, slicedRepos)
		cache = common.UpdateCache(cache, notificationPayload, similar)
		fmt.Println("\nnotificationPayload", notificationPayload)
		fmt.Println("similar", similar)
		fmt.Println("cache", cache)
	})
	c.Start()
	defer c.Stop()

	// Let the process runs forever
	a := make(chan os.Signal)
	signal.Notify(a, os.Interrupt, os.Kill)
	<-a
}

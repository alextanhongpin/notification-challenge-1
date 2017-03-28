package main

import (
	"fmt"
	"github.com/alextanhongpin/notification-challenge/common"
	"github.com/alextanhongpin/notification-challenge/model"
	"github.com/alextanhongpin/notification-challenge/service"
	"gopkg.in/robfig/cron.v2"
	"os"
	"os/signal"
	"sort"
)

var cache map[int]bool

func main() {
	fmt.Println("Started cron job")
	if cache == nil {
		cache = make(map[int]bool)
	}

	c := cron.New()
	// Run the cron job every thirty minutes
	c.AddFunc("0 */30 * * * *", func() {
		fmt.Println("Running cron job")

		repos, err := service.FetchPublicRepositories()
		if err != nil {
			fmt.Println(err)
		}
		slicedRepos := repos

		if len(slicedRepos) == 0 {
			return
		} else if len(slicedRepos) > 5 {
			slicedRepos = repos[0:5]
		}

		notificationPayload := common.MakeNotificationPayload(cache, slicedRepos)
		similar := common.GetSimilarData(cache, slicedRepos)
		cache = common.UpdateCache(cache, notificationPayload, similar)

		message := model.Message{
			Channel:     "#general", //"#intrvw-notification",
			Text:        "The last 5 updated repository (written in golang)",
			Username:    "alextanhongpin",
			IconEmoji:   ":ghost:",
			Attachments: []model.Attachment{},
		}

		for _, notification := range notificationPayload {
			// fmt.Println(int(notification.UpdatedAt.UnixNano() / 1000000))
			fmt.Println("Timestamp", notification.PushedAt.UnixNano()/1000000)
			message.Attachments = append(message.Attachments, model.Attachment{
				AuthorIcon: notification.Owner.AvatarURL,
				AuthorLink: notification.Owner.URL,
				AuthorName: notification.Owner.Login,
				Title:      notification.FullName,
				Footer:     notification.Description,
				TitleLink:  notification.HTMLURL,
				Timestamp:  notification.PushedAt.Unix(), // "epoch time"
			})
		}
		// if len(message.Attachments) == 0 {
		// 	message.Text = "There are no updates. Grab a :taco:!"
		// }

		fmt.Println("Sending...")

		sort.Slice(message.Attachments, func(i, j int) bool { return message.Attachments[i].Timestamp > message.Attachments[j].Timestamp })

		ok, err := service.PostToSlack(message)
		if err != nil {
			panic(err)
		}
		if ok {
			fmt.Println("Successfully send notification")
		}

		//fmt.Println("\nnotificationPayload", notificationPayload)
		//fmt.Println("similar", similar)
		//fmt.Println("cache", cache)
		fmt.Println("Sent", message)

	})
	c.Start()
	defer c.Stop()

	// Let the process runs forever
	a := make(chan os.Signal)
	signal.Notify(a, os.Interrupt, os.Kill)
	<-a
}

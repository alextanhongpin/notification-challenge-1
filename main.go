package main

import (
	"fmt"
	"github.com/alextanhongpin/notification-challenge/common"
	"github.com/alextanhongpin/notification-challenge/model"
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
	c.AddFunc("0 */1 * * * *", func() {
		fmt.Println("Running cron job")
		repos, err := service.FetchPublicRepositories()
		if err != nil {
			fmt.Println(err)
		}
		slicedRepos := repos[0:5]
		notificationPayload := common.MakeNotificationPayload(cache, slicedRepos)
		similar := common.GetSimilarData(cache, slicedRepos)
		cache = common.UpdateCache(cache, notificationPayload, similar)

		message := model.Message{
			Channel:     "#general",
			Text:        "The last 5 updated repository (golang)",
			Username:    "alextanhongpin",
			IconEmoji:   ":ghost:",
			Attachments: []model.Attachment{},
		}

		for _, notification := range notificationPayload {
			message.Attachments = append(message.Attachments, model.Attachment{
				AuthorIcon: notification.Owner.AvatarURL,
				AuthorLink: notification.Owner.URL,
				AuthorName: notification.Owner.Login,
				Title:      notification.FullName,
				Footer:     notification.Description,
				TitleLink:  notification.HTMLURL,
			})
		}
		if len(message.Attachments) == 0 {
			message.Text = "There are no updates. Grab a :taco:!"
		}
		service.PostToSlack(message)

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

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

type By func(r1, r2 *model.Attachment) bool

type attachmentsSorter struct {
	attachments []model.Attachment
	by          func(r1, r2 *model.Attachment) bool
}

func (by By) Sort(attachments []model.Attachment) {
	rs := &attachmentsSorter{
		attachments: attachments,
		by:          by,
	}
	sort.Sort(rs)
}

// Len is part of sort.Interface.
func (s *attachmentsSorter) Len() int {
	return len(s.attachments)
}

// Swap is part of sort.Interface.
func (s *attachmentsSorter) Swap(i, j int) {
	s.attachments[i], s.attachments[j] = s.attachments[j], s.attachments[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *attachmentsSorter) Less(i, j int) bool {
	return s.by(&s.attachments[i], &s.attachments[j])
}

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
		timestamp := func(r1, r2 *model.Attachment) bool {
			return r1.Timestamp > r2.Timestamp
		}
		fmt.Println("Sending...")

		By(timestamp).Sort(message.Attachments)

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

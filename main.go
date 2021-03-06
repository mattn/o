package main

import (
	"log"
	"regexp"
	"time"

	"github.com/mattn/go-gntp"
	rss "github.com/mattn/go-pkg-rss"
)

var (
	re = regexp.MustCompile(`^オッ\s*(https?://\S+)?`)
)

func main() {
	client := gntp.NewClient()
	client.AppName = "オッ"
	client.Register([]gntp.Notification{
		{Event: "オッ", Enabled: true},
	})
	feeder := rss.New(5, true, nil,
		func(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
			for _, item := range newitems {
				m := re.FindAllStringSubmatch(item.Title, -1)
				if len(m) == 0 {
					continue
				}
				log.Println(item.Title)
				client.Notify(&gntp.Message{
					Event:    "オッ",
					Title:    "オッRSS",
					Text:     item.Title,
					Icon:     "https://raw.githubusercontent.com/mattn/o/master/icon.png",
					Callback: m[0][1],
				})
			}
		},
	)
	for {
		err := feeder.Fetch(`https://queryfeed.net/twitter?q=from%3Atodesking&title-type=tweet-text-full`, nil)
		if err != nil {
			log.Print(err)
			continue
		}
		time.Sleep(time.Minute)
	}
}

package main

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/radovskyb/watcher"
)

const alertTemplate = "New file is ready 🎉\n\n%s"

func main() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	w := watcher.New()
	r := regexp.MustCompile(config.Filter)
	w.AddFilterHook(watcher.RegexFilterHook(r, false))
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Create, watcher.Move)

	go func() {
		for {
			select {
			case event := <-w.Event:
				message := fmt.Sprintf(alertTemplate, event.Name())
				SendTelegramMessage(message, config.Telegram.BotToken, config.Telegram.ChatID)
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	for i := range config.Directories {
		if err := w.Add(config.Directories[i]); err != nil {
			log.Fatalln(err)
		}
	}

	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}

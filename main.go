package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tidwall/pretty"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	prettyOptions := &pretty.Options{
		Indent: "    ",
		SortKeys: true,
	}

	b, err := tb.NewBot(tb.Settings{
		Token:  "",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/pretty", func(m *tb.Message) {
		fmt.Printf("/pretty: at %s from %s got: %s", m.Time().String(), m.Sender.Username, m.Payload)
		result := pretty.PrettyOptions([]byte(m.Payload), prettyOptions)
		b.Reply(m,  string(result))
	})

	b.Start()
}
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var arrayOfItems []Item

type Item struct {
	Name     string
	url      string
	ThumbURL string
	Type     string
	MIME     string
}

func logCommand(m *tb.Message, command string) {
	t := time.Now()
	log.Print(m.Sender.Username + " sent " + command + " at " + t.Format("2006-01-02 15:04:05"))
}

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select name,url,thumburl,type,mime from entries")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var url string
		var thumburl string
		var Type string
		var mime string
		err = rows.Scan(&name, &url, &thumburl, &Type, &mime)
		if err != nil {
			log.Fatal(err)
		}
		arrayOfItems = append(arrayOfItems, Item{Name: name, url: url, ThumbURL: thumburl, Type: Type, MIME: mime})

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	token := os.Getenv("TELEBOT_TOKEN")
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tb.OnQuery, func(q *tb.Query) {
		results := make(tb.Results, len(arrayOfItems)) // []tb.Result
		for i, item := range arrayOfItems {
			if strings.Contains(item.Type, "photo") {
				result := &tb.PhotoResult{
					URL:      item.url,
					ThumbURL: item.ThumbURL,
				}
				results[i] = result
				results[i].SetResultID(strconv.Itoa(i)) // It's needed to set a unique string ID for each result
			} else if strings.Contains(item.Type, "audio") {
				result := &tb.AudioResult{
					Title: item.Name,
					URL:   item.url,
				}
				results[i] = result
				results[i].SetResultID(strconv.Itoa(i)) // It's needed to set a unique string ID for each result
			} else if strings.Contains(item.Type, "document") {
				result := &tb.DocumentResult{
					Title: item.Name,
					URL:   item.url,
					MIME:  item.MIME,
				}
				results[i] = result
				results[i].SetResultID(strconv.Itoa(i)) // It's needed to set a unique string ID for each result
			}

		}

		err := b.Answer(q, &tb.QueryResponse{
			Results:   results,
			CacheTime: 60, // a minute
		})

		if err != nil {
			fmt.Println(err)
		}
	})

	log.Print("The bot is starting...")
	b.Start()
}

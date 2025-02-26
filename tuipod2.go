package main

import (
	"bufio"
	"encoding/xml"
	"fmt"

	//"io"
	//"math"
	//"net/http"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const app_name = "tuipod2"
const statusbar_template = "STATUS"

type Episode struct {
	Url   string
	Title string
}

type Podcast struct {
	Url      string
	Title    string
	Episodes []Episode
}

type Subscription struct {
	Text   string `xml:"text,attr"`
	XmlUrl string `xml:"xmlUrl,attr"`
}

type Enclosure struct {
	Url    string `xml:"url,attr"`
	Type   string `xml:"type,attr"`
	Length int64  `xml:"length,attr"`
}

type Item struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	PubDate     string    `xml:"pubDate"`
	Enclosure   Enclosure `xml:"enclosure"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Rss struct {
	Channel Channel `xml:"channel"`
}

func main() {
	//fmt.Println("tuipod2")

	// test data

	episode := Episode{"https://podcast.com/episode.mp3", "Sample Episode"}
	podcast := Podcast{"https://podcast.com/", "Sample Podcast", make([]Episode, 0)}

	podcast.Episodes = append(podcast.Episodes, episode)

	//fmt.Println("PODCAST:", podcast)

	// test opml

	file, err := os.Open("subscriptions.opml")

	if err != nil {
		fmt.Println("ERROR opening subscriptions.opml:", err)
	}

	defer file.Close()

	opml := make([]Subscription, 0)
	scanner := bufio.NewScanner((file))

	// TODO: get rid of entry == line expectation
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "xmlUrl") {
			subscription := &Subscription{}
			err = xml.Unmarshal([]byte(line), subscription)

			if err != nil {
				fmt.Println("ERROR extracting subscription detail:", err)
			}

			opml = append(opml, *subscription)
		}
	}

	//fmt.Println("OPML:", len(opml), opml)
	//fmt.Println("First Subscription", opml[0])

	// test network pull
	/*
		resp, err := http.Get(opml[0].XmlUrl)

		if resp != nil {
			fmt.Println("ERROR retrieving feeds:", err)
		}

		defer resp.Body.Close()
	*/

	//body, err := io.ReadAll(resp.Body)
	//body_as_string := string(body[:]) // TODO: tighten this up

	//fmt.Println("RESPONSE", body_as_string)

	// TODO: parse podcast feed into our data objects

	/*
		rss := Rss{}

		decoder := xml.NewDecoder(resp.Body)
		err = decoder.Decode(&rss)
		if err != nil {
			fmt.Println("ERROR parsing feeds:", err)
		}
	*/

	//fmt.Println(rss)

	app := tview.NewApplication()

	frame := tview.NewFrame(tview.NewBox().SetBackgroundColor(tcell.ColorBlue)).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(app_name, true, tview.AlignLeft, tcell.ColorWhite).
		AddText(statusbar_template, false, tview.AlignLeft, tcell.ColorWhite)

	// box := tview.NewBox().
	// 	SetBorder(true).
	// 	SetTitle(app_name)

	if err := app.SetRoot(frame, true).SetFocus(frame).Run(); err != nil {
		panic(err)
	}
}

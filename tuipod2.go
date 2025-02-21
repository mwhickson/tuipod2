package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

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

func main() {
	fmt.Println("tuipod2")

	episode := Episode{"https://podcast.com/episode.mp3", "Sample Episode"}
	podcast := Podcast{"https://podcast.com/", "Sample Podcast", make([]Episode, 0)}

	podcast.Episodes = append(podcast.Episodes, episode)

	fmt.Println("PODCAST:", podcast)

	file, err := os.Open("subscriptions.opml")

	if err != nil {
		fmt.Println("ERROR:", err)
	}

	defer file.Close()

	opml := make([]Subscription, 0)
	scanner := bufio.NewScanner((file))

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "xmlUrl") {
			subscription := &Subscription{}
			err = xml.Unmarshal([]byte(line), subscription)

			if err != nil {
				fmt.Println("ERROR:", err)
			}

			opml = append(opml, *subscription)
		}
	}

	fmt.Println("OPML:", len(opml), opml)
}

package tuipod2

import (
	"encoding/xml"
	"net/http"
)

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

type Feed struct {
	Channel Channel `xml:"channel"`
}

func RetrieveFeed(url string) Feed {
	resp, err := http.Get(url) // TODO: set client name

	if resp != nil {
		panic(err)
	}

	defer resp.Body.Close()

	feed := Feed{}

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&feed)
	if err != nil {
		panic(err)
	}

	return feed
}

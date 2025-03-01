package tuipod2

import (
	"encoding/xml"
	"os"
)

type Subscription struct {
	Text   string `xml:"text,attr"`
	XmlUrl string `xml:"xmlUrl,attr"`
}

type OutlineContainer struct {
	Items []Subscription `xml:"outline"`
}

type Body struct {
	Outline OutlineContainer `xml:"outline"`
}

type Opml struct {
	Document Body `xml:"body"`
}

func NewSubscription(url string, title string) *Subscription {
	s := &Subscription{Text: url, XmlUrl: title}
	return s
}

func LoadSubscriptions(subscription_file string) []Subscription {
	file, err := os.Open(subscription_file)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	opml := Opml{}

	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&opml)
	if err != nil {
		panic(err)
	}

	subscriptions := opml.Document.Outline.Items

	return subscriptions
}

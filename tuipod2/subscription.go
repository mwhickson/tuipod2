package tuipod2

import (
	"bufio"
	"encoding/xml"
	"os"
	"strings"
)

type Subscription struct {
	Text   string `xml:"text,attr"`
	XmlUrl string `xml:"xmlUrl,attr"`
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

	subscriptions := make([]Subscription, 0)
	scanner := bufio.NewScanner((file))

	// TODO: get rid of entry == line expectation
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "xmlUrl") {
			subscription := &Subscription{}
			err = xml.Unmarshal([]byte(line), subscription)

			if err != nil {
				panic(err)
			}

			subscriptions = append(subscriptions, *subscription)
		}
	}

	return subscriptions
}

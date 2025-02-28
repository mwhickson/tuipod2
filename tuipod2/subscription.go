package tuipod2

import (
	"bufio"
	"encoding/xml"
	"fmt"
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

func LoadSubscriptions() {
	file, err := os.Open("data/subscriptions.opml")

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

	// test opml

	//fmt.Println("OPML:", len(opml), opml)
	//fmt.Println("First Subscription", opml[0])
}

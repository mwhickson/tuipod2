package tuipod2

//"io"
//"net/http"

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

// func TestFeed() {
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
// }

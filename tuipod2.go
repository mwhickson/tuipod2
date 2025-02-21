package main

import (
	"bufio"
	"fmt"
	"os"
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

	opml := make([]string, 0)
	scanner := bufio.NewScanner((file))

	for scanner.Scan() {
		line := scanner.Text()
		opml = append(opml, line)
	}

	fmt.Println("OPML:", len(opml), opml)
}

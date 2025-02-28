package tests

import (
	"testing"

	tuipod2 "tuipod2/tuipod2"
)

func TestEpisode(t *testing.T) {
	episode := tuipod2.NewEpisode("https://podcast.com/episode.mp3", "Sample Episode")

	if episode.Title != "Sample Episode" {
		t.Fatalf(`Episode title not set correctly`)
	}

	if episode.Url != "https://podcast.com/episode.mp3" {
		t.Fatalf(`Episode url not set correctly`)
	}
}

func TestPodcast(t *testing.T) {
	episode := tuipod2.NewEpisode("https://podcast.com/episode.mp3", "Sample Episode")
	podcast := tuipod2.NewPodcast("https://podcast.com/", "Sample Podcast")
	podcast.Episodes = append(podcast.Episodes, *episode)

	if podcast.Title != "Sample Podcast" {
		t.Fatalf(`Podcast title not set correctly`)
	}

	if podcast.Url != "https://podcast.com/" {
		t.Fatalf(`Podcast url not set correctly`)
	}

	if len(podcast.Episodes) != 1 {
		t.Fatalf(`Podcast episodes not correct length`)
	}

	if podcast.Episodes[0].Url != "https://podcast.com/episode.mp3" {
		t.Fatalf(`Podcast episode url not set correctly`)
	}
}

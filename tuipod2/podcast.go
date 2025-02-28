package tuipod2

type Podcast struct {
	Url      string
	Title    string
	Episodes []Episode
}

func NewPodcast(url string, title string) *Podcast {
	p := &Podcast{Url: url, Title: title, Episodes: make([]Episode, 0)}
	return p
}

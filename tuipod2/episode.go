package tuipod2

type Episode struct {
	Url   string
	Title string
}

func NewEpisode(url string, title string) *Episode {
	e := &Episode{Url: url, Title: title}
	return e
}

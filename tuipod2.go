package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	//"io"
	"math"
	//"net/http"
	"os"
	"strings"

	// REF: https://pkg.go.dev/github.com/metafates/pat/color
	"github.com/charmbracelet/lipgloss"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"

	tea "github.com/charmbracelet/bubbletea"
)

const app_name = "tuipod2"
const statusbar_template = "STATUS"

type model struct {
	console_height int
	console_width  int
	search         textinput.Model
	podcast_table  table.Model
	episode_table  table.Model
}

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

func main() {
	//fmt.Println("tuipod2")

	// test data

	episode := Episode{"https://podcast.com/episode.mp3", "Sample Episode"}
	podcast := Podcast{"https://podcast.com/", "Sample Podcast", make([]Episode, 0)}

	podcast.Episodes = append(podcast.Episodes, episode)

	//fmt.Println("PODCAST:", podcast)

	// test opml

	file, err := os.Open("subscriptions.opml")

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

	//fmt.Println("OPML:", len(opml), opml)
	//fmt.Println("First Subscription", opml[0])

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

	// Bubble Tea UI
	p := tea.NewProgram(initialModel(opml), tea.WithAltScreen())
	if _, err = p.Run(); err != nil {
		fmt.Println("ERROR running program:", err)
		os.Exit(1)
	}
}

func initialModel(subscriptions []Subscription) model {
	ti := textinput.New()

	ti.Placeholder = "Search"
	ti.CharLimit = 255

	pt_columns := []table.Column{
		{Title: "Podcast Title", Width: 40},
	}

	et_columns := []table.Column{
		{Title: "Episode Title", Width: 40},
	}

	podcasts := make([]table.Row, 0)
	for _, subscription := range subscriptions {
		row := table.Row{subscription.Text}
		podcasts = append(podcasts, row)
	}

	pt := table.New(
		table.WithColumns(pt_columns),
		table.WithRows(podcasts),
	)

	pt.SetStyles(table.DefaultStyles())

	et := table.New(
		table.WithColumns(et_columns),
	)

	et.SetStyles(table.DefaultStyles())

	return model{
		console_height: 0,
		console_width:  0,
		search:         ti,
		podcast_table:  pt,
		episode_table:  et,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.console_height = msg.Height
		m.console_width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	m.search.Width = m.console_width
	m.search.Focus()

	m.search, cmd = m.search.Update(msg)

	return m, cmd
}

func (m model) View() string {
	s := ""
	s += RenderTitlebar(m) + "\n"
	s += RenderSearchbar(m) + "\n"
	s += RenderPodcastTable(m) + "\n"
	s += RenderEpisodeTable(m) + "\n"
	s += RenderStatusbar(m)

	return s
}

func RenderTitlebar(m model) string {
	titlebar := app_name
	titlebar_style := lipgloss.NewStyle().
		Width(m.console_width).
		Foreground(lipgloss.Color("11")).
		Background(lipgloss.Color("12"))

	s := titlebar_style.Render(titlebar)

	return s
}

func RenderSearchbar(m model) string {
	return m.search.View()
}

func RenderPodcastTable(m model) string {
	table_height := math.Ceil((float64(m.console_height) - 3.0) / 2.0)
	m.podcast_table.SetWidth(m.console_width) // height change works, but width doesn't seem to...
	m.podcast_table.SetHeight(int(table_height))
	return m.podcast_table.View()
}

func RenderEpisodeTable(m model) string {
	table_height := math.Floor((float64(m.console_height) - 3.0) / 2.0)
	m.episode_table.SetWidth(m.console_width) // height change works, but width doesn't seem to...
	m.episode_table.SetHeight(int(table_height))
	return m.episode_table.View()
}

func RenderStatusbar(m model) string {
	statusbar := statusbar_template
	statusbar_style := lipgloss.NewStyle().
		Width(m.console_width).
		Foreground(lipgloss.Color("8")).
		Background(lipgloss.Color("15"))

	s := statusbar_style.Render(statusbar)

	return s
}

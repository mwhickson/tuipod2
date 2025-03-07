package tuipod2

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const app_name = "tuipod2 v0.1"
const statusbar_template = "STATUS"

var app = tview.NewApplication()

var searchReference *tview.InputField

var podcastTableReference *tview.Table
var episodeTableReference *tview.Table
var pagesReference *tview.Pages

func RunApplication() {
	subscriptions := LoadSubscriptions("data/subscriptions.opml")

	search := makeSearch()
	podcast_table := makePodcastTable(subscriptions)
	episode_table := makeEpisodeTable()

	searchReference = search
	podcastTableReference = podcast_table
	episodeTableReference = episode_table

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(search, 1, 1, true).
		AddItem(podcast_table, 0, 1, false).
		AddItem(episode_table, 0, 1, false)

	frame := tview.NewFrame(flex).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(app_name, true, tview.AlignLeft, tcell.ColorWhite).
		AddText(statusbar_template, false, tview.AlignLeft, tcell.ColorWhite)

	pages := tview.NewPages().
		AddPage("main", frame, true, true)

	pagesReference = pages

	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}
}

func onSearchSubmitted(key tcell.Key) {
	if key == tcell.KeyTab {
		app.SetFocus(podcastTableReference)
	} else if key == tcell.KeyEnter {
		searchTerm := searchReference.GetText()
		onSearchExecute(searchTerm)
	}
}

func onSearchExecute(term string) {
	searchModal := makeSearchModal(term)
	pagesReference.AddPage("modal", searchModal, true, true)
}

func onCloseSearch() {
	pagesReference.RemovePage("modal")
	app.SetFocus(searchReference)
}

func makeSearch() *tview.InputField {
	search_field := tview.NewInputField().
		SetPlaceholder("search for a podcast...").
		SetDoneFunc(onSearchSubmitted)
	return search_field
}

func makeSearchModal(term string) tview.Primitive {
	form := tview.NewForm().
		AddInputField("Search", term, 20, nil, nil).
		AddButton("Ok", onCloseSearch)

	form.SetBorder(true).SetTitle("Search")

	return makeModal(form, 40, 10)
}

func makeModal(p tview.Primitive, width int, height int) tview.Primitive {
	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(
			tview.NewFlex().
				SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, height, 1, true).
				AddItem(nil, 0, 1, false), width, 1, true,
		).AddItem(nil, 0, 1, false)
}

func onPodcastTableDone(key tcell.Key) {
	if key == tcell.KeyTab {
		app.SetFocus(episodeTableReference)
	}
}

func makePodcastTable(subscriptions []Subscription) *tview.Table {
	podcast_table := tview.NewTable()

	podcast_table.SetTitle("Podcasts")
	podcast_table.SetBorder(true)
	podcast_table.SetBackgroundColor(tcell.ColorDarkBlue)
	podcast_table.SetSelectable(true, false)
	podcast_table.SetDoneFunc(onPodcastTableDone)

	for i := range subscriptions {
		podcast_table.SetCell(i, 0, tview.NewTableCell(subscriptions[i].Text))
	}

	return podcast_table
}

func onEpisodeTableDone(key tcell.Key) {
	if key == tcell.KeyTab {
		app.SetFocus(searchReference)
	}
}

func makeEpisodeTable() *tview.Table {
	episode_table := tview.NewTable()

	episode_table.SetTitle("Episodes")
	episode_table.SetBorder(true)
	episode_table.SetBackgroundColor(tcell.ColorBlue)
	episode_table.SetSelectable(true, false)
	episode_table.SetDoneFunc(onEpisodeTableDone)

	return episode_table
}

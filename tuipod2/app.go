package tuipod2

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const app_name = "tuipod2 v0.1"
const statusbar_template = "STATUS"

func RunApplication() {
	app := tview.NewApplication()

	subscriptions := LoadSubscriptions("data/subscriptions.opml")

	search := makeSearch()
	podcast_table := makePodcastTable(subscriptions)
	episode_table := makeEpisodeTable()

	// bring it all together

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(search, 1, 1, true).
		AddItem(podcast_table, 0, 1, false).
		AddItem(episode_table, 0, 1, false)

	frame := tview.NewFrame(flex).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(app_name, true, tview.AlignLeft, tcell.ColorWhite).
		AddText(statusbar_template, false, tview.AlignLeft, tcell.ColorWhite)

	// TODO: figure out focus switching...
	// if err := app.SetRoot(frame, true).SetFocus(frame).Run(); err != nil {
	if err := app.SetRoot(frame, true).SetFocus(podcast_table).Run(); err != nil {
		panic(err)
	}
}

func makeSearch() *tview.InputField {
	return tview.NewInputField().SetPlaceholder("search for a podcast...")
}

func makePodcastTable(subscriptions []Subscription) *tview.Table {
	podcast_table := tview.NewTable()

	podcast_table.SetTitle("Podcasts")
	podcast_table.SetBorder(true)
	podcast_table.SetBackgroundColor(tcell.ColorDarkBlue)
	podcast_table.SetSelectable(true, false)

	for i := range subscriptions {
		podcast_table.SetCell(i, 0, tview.NewTableCell(subscriptions[i].Text))
	}

	return podcast_table
}

func makeEpisodeTable() *tview.Table {
	episode_table := tview.NewTable()

	episode_table.SetTitle("Episodes")
	episode_table.SetBorder(true)
	episode_table.SetBackgroundColor(tcell.ColorBlue)
	episode_table.SetSelectable(true, false)

	return episode_table
}

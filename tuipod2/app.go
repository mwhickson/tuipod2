package tuipod2

import (
	"github.com/rivo/tview"
)

const app_name = "tuipod2 v0.1"
const statusbar_template = "STATUS"

func RunApplication() {
	app := tview.NewApplication()

	search := tview.NewInputField().SetPlaceholder("search for a podcast...")

	// set up podcast table

	podcast_table := tview.NewTable()

	podcast_table.SetTitle("Podcasts")
	podcast_table.SetBorder(true)
	// podcast_table.SetBackgroundColor(tcell.ColorDarkBlue)
	podcast_table.SetSelectable(true, false)

	// for i := range opml {
	// 	podcast_table.SetCell(i, 0, tview.NewTableCell(opml[i].Text))
	// }

	// set up episodes table

	episode_table := tview.NewTable()

	episode_table.SetTitle("Episodes")
	episode_table.SetBorder(true)
	// episode_table.SetBackgroundColor(tcell.ColorBlue)

	// bring it all together

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(search, 1, 1, true).
		AddItem(podcast_table, 0, 1, false).
		AddItem(episode_table, 0, 1, false)

	frame := tview.NewFrame(flex).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(app_name, true, tview.AlignLeft, 0).           //, tcell.ColorWhite).
		AddText(statusbar_template, false, tview.AlignLeft, 0) //, tcell.ColorWhite)

	//if err := app.SetRoot(frame, true).SetFocus(frame).Run(); err != nil {
	if err := app.SetRoot(frame, true).SetFocus(podcast_table).Run(); err != nil {
		panic(err)
	}
}

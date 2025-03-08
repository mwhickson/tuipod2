package tuipod2

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const app_name = "tuipod2 v0.1"
const statusbar_template = "STATUS | Search (Ctrl+S) | Quit (ESC)"

var appReference *tview.Application

// var searchReference *tview.InputField

var podcastTableReference *tview.Table
var episodeTableReference *tview.Table
var pagesReference *tview.Pages

func RunApplication() {
	app := tview.NewApplication()

	appReference = app
	appReference.SetInputCapture(onAppInputCapture)

	subscriptions := LoadSubscriptions("data/subscriptions.opml")

	podcast_table := makePodcastTable(subscriptions)
	episode_table := makeEpisodeTable()

	podcastTableReference = podcast_table
	episodeTableReference = episode_table

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(podcast_table, 0, 1, false).
		AddItem(episode_table, 0, 1, false)

	frame := tview.NewFrame(flex).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(app_name, true, tview.AlignLeft, tcell.ColorWhite).
		AddText(statusbar_template, false, tview.AlignLeft, tcell.ColorWhite)

	pages := tview.NewPages().
		AddPage("main", frame, true, true)

	pagesReference = pages

	if err := app.SetRoot(pages, true).SetFocus(podcast_table).Run(); err != nil {
		panic(err)
	}
}

func onAppInputCapture(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyEscape {
		onQuitConfirmExecute()
		return nil
	} else if event.Key() == tcell.KeyCtrlS {
		onSearchExecute("")
		return nil
	}

	return event
}

func onQuitConfirmExecute() {
	quitConfirmModal := makeQuitConfirmModal()
	pagesReference.AddPage("quitconfirm", quitConfirmModal, true, true)
}

func onCancelQuit() {
	pagesReference.RemovePage("quitconfirm")
}

func onSearchExecute(term string) {
	searchModal := makeSearchModal(term)
	pagesReference.AddPage("searchform", searchModal, true, true)
}

func onCloseSearch() {
	pagesReference.RemovePage("searchform")
}

func makeQuitConfirmModal() tview.Primitive {
	quitModal := tview.NewModal()

	quitModal.SetTitle("Quit")
	quitModal.SetBorder(true)
	quitModal.SetText("Really quit?")
	quitModal.AddButtons([]string{"Ok", "Cancel"})
	quitModal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "Ok" {
			appReference.Stop()
		} else {
			onCancelQuit()
		}
	})

	return quitModal
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
		appReference.SetFocus(episodeTableReference)
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
		appReference.SetFocus(podcastTableReference)
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

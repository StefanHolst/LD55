package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type WelcomeView struct {
	tview.Flex
}

func NewWelcomeView() *WelcomeView {
	welcomeView := &WelcomeView{
		Flex: *tview.NewFlex().SetDirection(tview.FlexRow),
	}

	welcomeView.
		AddItem(tview.NewTextView().SetText("Welcome to the game!").SetTextAlign(tview.AlignCenter), 0, 4, false).
		AddItem(tview.NewTextView().
			SetText("You have found an old secret recipe for summoning valuable items. You have to build a business around it.").
			SetTextAlign(tview.AlignCenter), 0, 1, false).
		AddItem(tview.NewTextView().
			SetText("But beware of the jealous people in the village, if they find out about you secret. They will kill you.").
			SetTextAlign(tview.AlignCenter), 0, 1, false).
		AddItem(tview.NewTextView().
			SetText("The more you Sell the more people will notice it. If enough people notice it, you will be killed. If you don't Sell enough the demon will kill you.").
			SetTextAlign(tview.AlignCenter), 0, 1, false)

	welcomeView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			pages.SwitchToPage("game")
		}

		return event
	})

	return welcomeView
}

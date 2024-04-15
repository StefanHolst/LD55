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
		AddItem(tview.NewTextView().SetText("The Travelling Salesman's Mission!").SetTextAlign(tview.AlignCenter), 0, 4, false).
		AddItem(tview.NewTextView().
			SetText("You are a salesman travelling the world to sell goods. Your mission is to collect all the collectable artifacts.").
			SetTextAlign(tview.AlignCenter), 0, 1, false)

	welcomeView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			pages.SwitchToPage("game")
		}

		return event
	})

	return welcomeView
}

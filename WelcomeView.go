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
			SetText("In the vibrant city of Shagadelic, Alexander Powers, a seasoned travelling salesman with a knack for adventure, stumbled upon a secretive society known as the Collectors' Guild. Intrigued by their mission to acquire rare artifacts, Alexander joined their ranks.").
			SetTextAlign(tview.AlignCenter), 0, 1, false).
		AddItem(tview.NewTextView().
			SetText("Tasked with gathering legendary relics scattered across distant lands, Alexander embarked on a perilous journey. He acquired the Horn of Frau Farbissina from the icy fjords of Nordheim and the Sakura Kimono from the mystical shrines of Dr. Evil's Volcano Lair.").
			SetTextAlign(tview.AlignCenter), 0, 1, false).
		AddItem(tview.NewTextView().
			SetText("Yet, shadows lurked in Shagadelic. A rival faction, led by a sinister figure known as Dougie, coveted the artifacts for malevolent purposes. One stormy night, Alexander was ambushed, losing the Statues Compass to these nefarious foes.").
			SetTextAlign(tview.AlignCenter), 0, 1, false).
		AddItem(tview.NewTextView().
			SetText("Determined to reclaim what was lost, Alexander pursued the shadowy figures to their hidden lair beneath the city. In a climactic battle of wits and wills, Alexander bested their enigmatic leader, retrieving the stolen artifact.").
			SetTextAlign(tview.AlignCenter), 0, 1, false).
		AddItem(tview.NewTextView().
			SetText("Returning triumphant to the guild, Alexander shared tales of his harrowing adventures. As the artifacts were gathered, the Collectors' Guild celebrated their triumph over adversity, knowing that their quest for treasures and mysteries had only just begun.").
			SetTextAlign(tview.AlignCenter), 0, 1, false)

	welcomeView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			pages.SwitchToPage("game")
		}

		return event
	})

	return welcomeView
}

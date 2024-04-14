package main

import (
	"github.com/rivo/tview"
	"strconv"
)

type PlayerInventoryView struct {
	*tview.Grid

	Player      *Player
	primitives  map[string]tview.Primitive
	accountView *tview.TextView
}

func NewPlayerInventoryView(player *Player) *PlayerInventoryView {
	p := &PlayerInventoryView{
		Grid:       tview.NewGrid(),
		Player:     player,
		primitives: map[string]tview.Primitive{},
	}
	p.SetRows(2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1).SetColumns(0, 0)

	p.AddItem(tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("Inventory"), 0, 0, 1, 2, 0, 0, false)

	// inventory
	for i := 0; i < 10; i++ {
		// product name
		p.AddItem(tview.NewTextView().SetText(ProductTypeToString[ProductType(i)]), i+1, 0, 1, 1, 0, 0, false)

		// amount
		p.primitives[ProductTypeToString[ProductType(i)]] = tview.NewTextView()
		p.AddItem(p.primitives[ProductTypeToString[ProductType(i)]], i+1, 1, 1, 1, 0, 0, false)
	}

	return p
}

func (p *PlayerInventoryView) UpdateViews() {
	for i := 0; i < len(p.Player.Inventory); i++ {
		p.primitives[ProductTypeToString[ProductType(i)]].(*tview.TextView).SetText(strconv.Itoa(len(p.Player.Inventory[ProductType(i)].Items)))
	}
}

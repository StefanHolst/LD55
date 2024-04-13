package main

import (
	"fmt"
	"github.com/rivo/tview"
	"strconv"
)

type InventoryView struct {
	*tview.Grid

	tableItems [][]tview.Primitive
	City       *City
	Player     *Player
}

func NewInventoryView(city *City, player *Player) *InventoryView {
	inv := &InventoryView{
		Grid:   tview.NewGrid().SetBorders(true),
		City:   city,
		Player: player,
	}

	inv.tableItems = make([][]tview.Primitive, 10)
	if player != nil {
		for i := 0; i < len(city.TradingItems); i++ {
			inv.tableItems[i] = make([]tview.Primitive, 4)
		}
	} else {
		for i := 0; i < len(city.TradingItems); i++ {
			inv.tableItems[i] = make([]tview.Primitive, 3)
		}
	}

	inv.setupTable()

	return inv
}

func (inv *InventoryView) setupTable() {
	if inv.Player == nil {
		inv.SetRows(make([]int, 12)...).SetColumns(make([]int, 3)...)
	} else {
		inv.SetColumns(make([]int, 4)...)
	}

	//inv.AddItem(tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(city.Name), 0, 0, 1, 4, 0, 0, false)

	rowOffset := 2

	// Column headers
	inv.AddItem(tview.NewTextView().SetText("Product"), rowOffset-1, 0, 1, 1, 0, 0, false)
	inv.AddItem(tview.NewTextView().SetText("Amount"), rowOffset-1, 1, 1, 1, 0, 0, false)
	inv.AddItem(tview.NewTextView().SetText("Price"), rowOffset-1, 2, 1, 1, 0, 0, false)
	if inv.Player != nil {
		inv.AddItem(tview.NewTextView().SetText("Player Amount"), rowOffset-1, 3, 1, 1, 0, 0, false)
	}

	// Table items
	for i := 0; i < len(inv.City.TradingItems); i++ {
		key := Products(i)

		// product name
		inv.tableItems[i][0] = tview.NewTextView().SetText(ProductsToString[key])
		inv.AddItem(inv.tableItems[i][0], i+rowOffset, 0, 1, 1, 0, 0, false)

		// amount
		inv.tableItems[i][1] = tview.NewTextView().SetText(strconv.Itoa(inv.City.Production[key]))
		inv.AddItem(inv.tableItems[i][1], i+rowOffset, 1, 1, 1, 0, 0, false)

		// Buttons and price
		cell := tview.NewGrid().SetRows(0).SetColumns(0, 0, 0)
		inv.AddItem(cell, i+rowOffset, 2, 1, 1, 0, 0, false)
		// price
		inv.tableItems[i][2] = tview.NewTextView().SetText(fmt.Sprintf("%f", inv.City.TradingItems[key].Value))
		cell.AddItem(inv.tableItems[i][2], 0, 0, 1, 1, 0, 0, false)
		// buy button
		button := tview.NewButton("Buy").SetSelectedFunc(func() { handleButton(true, inv.City, inv.Player, key) })
		cell.AddItem(button, 0, 1, 1, 1, 0, 0, false)
		// sell button
		button = tview.NewButton("Buy").SetSelectedFunc(func() { handleButton(false, inv.City, inv.Player, key) })
		cell.AddItem(button, 0, 2, 1, 1, 0, 0, false)

		// player amount
		if inv.Player != nil {
			inv.tableItems[i][3] = tview.NewTextView().SetText(strconv.Itoa(inv.Player.Inventory[key].Amount))
			inv.AddItem(inv.tableItems[i][3], i+rowOffset, 3, 1, 1, 0, 0, false)
		}
	}
}

func (inv *InventoryView) UpdateTable() {
	for i := 0; i < len(inv.City.TradingItems); i++ {
		key := Products(i)

		// amount
		inv.tableItems[i][1].(*tview.TextView).SetText(strconv.Itoa(inv.City.TradingItems[key].Amount))

		// price
		inv.tableItems[i][2].(*tview.TextView).SetText(fmt.Sprintf("%f", inv.City.TradingItems[key].Value))

		// player amount
		if inv.Player != nil {
			inv.tableItems[i][3].(*tview.TextView).SetText(strconv.Itoa(inv.Player.Inventory[key].Amount))
		}
	}
}

func handleButton(isBuy bool, city *City, player *Player, product Products) {
	if isBuy {
		player.buy(city, product)
	} else {
		player.sell(city, product)
	}

	inventoryView.UpdateTable()
}

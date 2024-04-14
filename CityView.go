package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"math/rand"
	"strconv"
)

type CityView struct {
	*tview.Grid

	primitives map[string]tview.Primitive

	City *City
}

func NewCityView(city *City) *CityView {
	inv := &CityView{
		Grid:       tview.NewGrid().SetBorders(true),
		primitives: map[string]tview.Primitive{},
		City:       city,
	}

	inv.SetInputCapture(inv.handleButtonEvents)

	inv.setupTable()

	return inv
}

func (cv *CityView) setupTable() {
	cv.SetRows(1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1).SetColumns(make([]int, 4)...)

	// City name
	cv.primitives["City-name"] = tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(cv.City.Name)
	cv.AddItem(cv.primitives["City-name"], 0, 0, 1, 4, 0, 0, false)

	rowOffset := 2

	// Column headers
	cv.AddItem(tview.NewTextView().SetText("Product"), rowOffset-1, 0, 1, 1, 0, 0, false)
	cv.AddItem(tview.NewTextView().SetText("Amount"), rowOffset-1, 1, 1, 1, 0, 0, false)
	cv.AddItem(tview.NewTextView().SetText("Buying Price"), rowOffset-1, 2, 1, 1, 0, 0, false)
	cv.AddItem(tview.NewTextView().SetText("Selling Price"), rowOffset-1, 3, 1, 1, 0, 0, false)

	// Table items
	for i := 0; i < len(cv.City.Inventory); i++ {
		key := ProductType(i)

		// product name
		cv.AddItem(tview.NewTextView().SetText(ProductTypeToString[key]), i+rowOffset, 0, 1, 1, 0, 0, false)

		// amount
		cv.primitives[ProductTypeToString[key]+"amount"] = tview.NewTextView()
		cv.AddItem(cv.primitives[ProductTypeToString[key]+"amount"], i+rowOffset, 1, 1, 1, 0, 0, false)

		// Buttons and price
		cell := tview.NewGrid().SetRows(0).SetColumns(0, 0)
		cv.AddItem(cell, i+rowOffset, 2, 1, 1, 0, 0, false)

		// buy price
		cv.primitives[ProductTypeToString[key]+"buy-price"] = tview.NewTextView()
		cell.AddItem(cv.primitives[ProductTypeToString[key]+"buy-price"], 0, 0, 1, 1, 0, 0, false)
		// Buy button
		cell.AddItem(tview.NewButton("Buy").SetSelectedFunc(func() { cv.handleButton(true, cv.City, key) }), 0, 1, 1, 1, 0, 0, false)

		cell = tview.NewGrid().SetRows(0).SetColumns(0, 0)
		cv.AddItem(cell, i+rowOffset, 3, 1, 1, 0, 0, false)

		// sell price
		cv.primitives[ProductTypeToString[key]+"sell-price"] = tview.NewTextView()
		cell.AddItem(cv.primitives[ProductTypeToString[key]+"sell-price"], 0, 0, 1, 1, 0, 0, false)
		// Sell button
		cell.AddItem(tview.NewButton("Sell").SetSelectedFunc(func() { cv.handleButton(false, cv.City, key) }), 0, 1, 1, 1, 0, 0, false)
	}

	// Artifacts
	cv.primitives["Artifact"] = tview.NewTextView()
	cv.AddItem(cv.primitives["Artifact"], rowOffset+12, 0, 1, 1, 0, 0, false)

	// Artifact value
	cv.primitives["Artifact-value"] = tview.NewTextView()
	cv.AddItem(cv.primitives["Artifact-value"], rowOffset+12, 1, 1, 1, 0, 0, false)
	cv.AddItem(tview.NewButton("Buy").SetSelectedFunc(func() {
		if player.BuyArtifact(cv.City) == false {
			return
		}

		// set artifact as sold
		cv.City.ArtifactSold = true

		// update artifacts prices
		for _, city := range gameLayout.Cities {
			city.Artifact.Value *= rand.Float64() + 1
		}

		gameLayout.UpdateViews()

	}), rowOffset+12, 2, 1, 1, 0, 0, false)
}

func (cv *CityView) UpdateViews() {
	// Update City name
	cv.primitives["City-name"].(*tview.TextView).SetText(cv.City.Name)

	for i := 0; i < len(cv.City.Inventory); i++ {
		key := ProductType(i)

		// amount
		cv.primitives[ProductTypeToString[key]+"amount"].(*tview.TextView).SetText(strconv.Itoa(cv.City.Inventory[key].Amount))

		// buy price
		cv.primitives[ProductTypeToString[key]+"buy-price"].(*tview.TextView).SetText(fmt.Sprintf("%f", cv.City.Inventory[key].Value))

		// sell price
		cv.primitives[ProductTypeToString[key]+"sell-price"].(*tview.TextView).SetText(fmt.Sprintf("%f", cv.City.Inventory[key].Value*0.8))
	}
}

func (cv *CityView) handleButton(isBuy bool, city *City, productType ProductType) {
	if isBuy {
		player.Buy(city, productType)
	} else {
		player.Sell(city, productType)
	}

	gameLayout.UpdateViews()
}

func (cv *CityView) handleButtonEvents(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyEsc {
		gameLayout.pages.SwitchToPage("world")
	}
	return event
}

func (cv *CityView) updateCity(city *City) {
	cv.City = city
	cv.UpdateViews()

	// Update Artifact
	cv.primitives["Artifact"].(*tview.TextView).SetText(ArtifactTypeToString[cv.City.Artifact.ArtifactType])

	// Update Artifact value
	if cv.City.ArtifactSold {
		cv.primitives["Artifact-value"].(*tview.TextView).SetText("Sold out")
	} else {
		cv.primitives["Artifact-value"].(*tview.TextView).SetText(fmt.Sprintf("%f", cv.City.Artifact.Value))
	}
}

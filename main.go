package main

import (
	"github.com/rivo/tview"
)

var running bool
var app *tview.Application
var grid *tview.Grid // Main grid
var inventoryView *InventoryView
var cities []*City
var player *Player

func main() {
	// Setup game logic
	cities = append(cities, NewCity("Aalborg", 1000, 81, 16))
	cities = append(cities, NewCity("Tokyo", 1000, 136, 21))
	cities = append(cities, NewCity("New York", 1000, 47, 19))
	cities = append(cities, NewCity("Las Vegas", 1000, 26, 21))
	cities = append(cities, NewCity("Rio de Janeiro", 1000, 57, 37))
	cities = append(cities, NewCity("Cape Town", 1000, 86, 40))
	cities = append(cities, NewCity("Moscow", 1000, 96, 17))
	cities = append(cities, NewCity("Sydney", 1000, 141, 40))
	player = NewPlayer("Player")

	// Create application
	app = tview.NewApplication()
	pages := tview.NewPages()

	grid = tview.NewGrid().
		SetRows(0).
		SetColumns(0).
		SetBorders(true)

	//inventoryView = NewInventoryView(city, player)
	//grid.AddItem(inventoryView, 0, 0, 1, 1, 0, 0, false)

	world := NewMapView(cities)
	grid.AddItem(world, 0, 0, 1, 1, 0, 0, false)

	//go mainLoop()

	grid.
		//AddItem(world, 0, 0, 1, 1, 0, 0, false).
		SetMouseCapture(handleMouseEvents).
		SetInputCapture(handleButtonEvents)

	// Run the application and handle any errors that occur.
	running = true
	app.SetRoot(grid, true).
		EnableMouse(true)
	if err := app.Run(); err != nil {
		running = false
		panic(err)
	}
	running = false
}

func mainLoop() {
	// do something
	//for running {
	//	city.UpdateProduction()
	//	city.UpdateUsage()
	//
	//	app.QueueUpdateDraw(inventoryView.UpdateTable)
	//
	//	time.Sleep(1 * time.Second)
	//}
}

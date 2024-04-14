package main

import (
	"github.com/rivo/tview"
)

var running bool

var app *tview.Application
var pages *tview.Pages
var gameLayout *GameLayoutView

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

	gameLayout = NewGameLayoutView(player, cities)

	// Create application
	app = tview.NewApplication()
	pages = tview.NewPages().
		AddPage("welcome", NewWelcomeView(), true, true).
		AddPage("game", gameLayout, true, false)

	// Run the application and handle any errors that occur.
	running = true
	app.SetRoot(pages, true).
		SetFocus(pages).
		EnableMouse(true)
	if err := app.Run(); err != nil {
		running = false
		panic(err)
	}
	running = false
}

func showDialog(message string) {
	modal := tview.NewModal()
	modal.
		SetText(message).
		AddButtons([]string{"Ok"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			app.SetRoot(pages, true).SetFocus(pages)
		})

	if err := app.SetRoot(modal, false).SetFocus(modal).Run(); err != nil {
		panic(err)
	}
}

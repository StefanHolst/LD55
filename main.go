package main

import (
	"embed"
	"github.com/rivo/tview"
	"io"
	"os"
	"path"
)

//go:embed payload/*
var payload embed.FS

var running bool

var app *tview.Application
var pages *tview.Pages
var gameLayout *GameLayoutView

var cities []*City
var player *Player

func main() {
	// unpack payload
	files, _ := payload.ReadDir("payload")
	for _, file := range files {
		sourceFileName := path.Join("payload", file.Name())
		destFileName := path.Join(".", file.Name())
		destFile, _ := os.Create(destFileName)
		defer destFile.Close()

		// open payload stream
		stream, _ := payload.Open(sourceFileName)
		// Copy payload stream to file
		_, _ = io.Copy(destFile, stream)
	}

	// Setup game logic
	cities = append(cities, NewCity("Aalborg", 1000, 79, 15, NewArtifact(TheHornOfFreyja)))
	cities = append(cities, NewCity("Tokyo", 1000, 133, 21, NewArtifact(TheSakuraKimono)))
	cities = append(cities, NewCity("New York", 1000, 46, 19, NewArtifact(TheStatuesCompass)))
	cities = append(cities, NewCity("Las Vegas", 1000, 25, 21, NewArtifact(TheLuckyChip)))
	cities = append(cities, NewCity("Rio de Janeiro", 1000, 57, 36, NewArtifact(TheCarnivalMask)))
	cities = append(cities, NewCity("Cape Town", 1000, 85, 38, NewArtifact(TheDiamondKey)))
	cities = append(cities, NewCity("Moscow", 1000, 95, 16, NewArtifact(TheWinterCloak)))
	cities = append(cities, NewCity("Sydney", 1000, 140, 39, NewArtifact(TheDreamtimeBoomerang)))
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

func showWin() {
	// Add win to pages
	// Add text to pages

	modal := func(p tview.Primitive, width, height int) tview.Primitive {
		return tview.NewGrid().
			SetColumns(0, width, 0).
			SetRows(0, height, 0).
			AddItem(p, 1, 1, 1, 1, 0, 0, true)
	}

	winningText := "Muhuhua!!\n\nYou gathered all my artifacts! I am free!"

	box := tview.NewGrid().SetColumns(0).SetRows(0).AddItem(tview.NewTextView().SetText(winningText).SetTextAlign(tview.AlignCenter), 0, 0, 1, 1, 0, 0, false).SetBorders(true)

	//box := tview.NewTextView().SetText("hej").
	//	SetBorder(true).
	//	SetTitle("You won")

	pages.HidePage("game")

	pages.AddPage("background", NewWinView(), true, true).
		AddPage("modal", modal(box, 40, 10), true, true)
}

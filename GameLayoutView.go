package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"time"
)

type GameLayoutView struct {
	*tview.Grid

	// left side
	pages     *tview.Pages
	worldView *MapView
	cityView  *CityView

	// Right side
	rightViewFlex       *tview.Flex
	playerStatsView     *PlayerStatsView
	playerInventoryView *PlayerInventoryView

	// Player
	Player *Player

	// Current City
	City *City

	// Cities
	Cities []*City

	// Time
	time time.Time
}

func NewGameLayoutView(player *Player, cities []*City) *GameLayoutView {
	glv := &GameLayoutView{
		Grid:   tview.NewGrid().SetBorders(true),
		Player: player,
		City:   cities[0],
		Cities: cities,
	}

	glv.time, _ = time.Parse("15:04:05", "00:00:00")

	// world
	glv.worldView = NewMapView(cities, player)

	// Cities
	glv.cityView = NewCityView(cities[0])

	// make pages
	glv.pages = tview.NewPages().
		AddPage("world", glv.worldView, true, true).
		AddPage("City", glv.cityView, true, false)

	// Player views
	glv.playerStatsView = NewPlayerStatsView(player)
	glv.playerInventoryView = NewPlayerInventoryView(player)

	glv.rightViewFlex = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(glv.playerStatsView, 0, 1, false).
		AddItem(glv.playerInventoryView, 0, 1, false)

	// setup grid
	glv.SetInputCapture(glv.handleButtonEvents)
	glv.SetRows(0).SetColumns(0, 40)
	glv.AddItem(glv.pages, 0, 0, 1, 1, 0, 0, false)
	glv.AddItem(glv.rightViewFlex, 0, 1, 1, 1, 0, 0, false)

	go glv.logicLoop()
	go glv.mainLoop()

	return glv
}

func (glv *GameLayoutView) handleButtonEvents(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyEsc {
		page, _ := glv.pages.GetFrontPage()
		if page == "City" {
			glv.pages.SwitchToPage("world")
		}
	}
	return event
}

func (glv *GameLayoutView) logicLoop() {
	for running {
		app.QueueUpdate(func() {
			for _, city := range cities {
				city.UpdateProduction()
				city.UpdateUsage()
			}
		})

		time.Sleep(1 * time.Second)
	}
}

var currentDay int

func (glv *GameLayoutView) mainLoop() {
	for running {
		glv.time = glv.time.Add(1 * time.Minute)
		// If the player is travelling, make the time go faster
		if glv.Player.IsTravelling {
			glv.time = glv.time.Add(15 * time.Minute)
		}

		// When we are changed day,
		if glv.time.Day() != currentDay {
			currentDay = glv.time.Day()
			for _, city := range cities {
				city.UpdateProductionScale()
				city.UpdateUsageScale()
			}
		}

		app.QueueUpdateDraw(glv.UpdateViews)

		time.Sleep(100 * time.Millisecond)
	}
}

func (glv *GameLayoutView) UpdateViews() {
	glv.cityView.UpdateViews()
	glv.playerInventoryView.UpdateViews()
	glv.playerStatsView.UpdateViews()
}

func (glv *GameLayoutView) UpdateCity(city *City) {
	glv.City = city
	glv.cityView.updateCity(city)
	glv.UpdateViews()
}

func (glv *GameLayoutView) ChangePage(page string) {
	glv.pages.SwitchToPage(page)
}

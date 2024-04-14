package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
)

type MapView struct {
	*tview.Box

	mapImage string

	Cities []*City

	fromCity *City
	toCity   *City

	xSlopeIntercept func(int) int
	ySlopeIntercept func(int) int
}

func NewMapView(cities []*City) *MapView {
	content, _ := os.ReadFile("drawnMap.txt")

	mapView := &MapView{
		Box:      tview.NewBox(),
		mapImage: string(content),
		Cities:   cities,
	}

	mapView.SetMouseCapture(mapView.handleMouseEvents)

	return mapView
}

func (m *MapView) Draw(screen tcell.Screen) {
	m.Box.DrawForSubclass(screen, m)

	// Draw the map
	x, y, i, imageLen := 0, 0, 0, len(m.mapImage)
	for true {
		if i >= imageLen {
			break
		}
		if m.mapImage[i] == '\n' {
			x = 0
			y++
			i++
		}
		tview.Print(screen, string(m.mapImage[i]), x, y, 1, tview.AlignCenter, tcell.ColorWhite)
		x++
		i++
	}

	// Draw the Cities
	for _, city := range m.Cities {
		color := tcell.ColorRed
		if city == gameLayout.cityView.City {
			color = tcell.ColorBlue
		}

		tview.Print(screen, "\u256d", city.X, city.Y, 1, tview.AlignCenter, color)
		tview.Print(screen, "\u2500", city.X+1, city.Y, 1, tview.AlignCenter, color)
		tview.Print(screen, "\u256e", city.X+2, city.Y, 1, tview.AlignCenter, color)
		tview.Print(screen, "\u2570", city.X, city.Y+1, 1, tview.AlignCenter, color)
		tview.Print(screen, "\u2500", city.X+1, city.Y+1, 1, tview.AlignCenter, color)
		tview.Print(screen, "\u256f", city.X+2, city.Y+1, 1, tview.AlignCenter, color)
	}

	// Draw the Player

	// Draw the path to the City

	if m.fromCity != nil && m.toCity != nil {
		// Draw line between cities
		x1 := m.fromCity.X
		x2 := m.toCity.X
		if m.fromCity.X > m.toCity.X {
			x1 = m.toCity.X
			x2 = m.fromCity.X
		}

		for x := x1; x < x2; x++ {
			y := m.ySlopeIntercept(x)
			tview.Print(screen, "#", x+1, y+1, 1, tview.AlignCenter, tcell.ColorGreen)
		}

		y1 := m.fromCity.Y
		y2 := m.toCity.Y
		if m.fromCity.Y > m.toCity.Y {
			y1 = m.toCity.Y
			y2 = m.fromCity.Y
		}

		for y := y1; y < y2; y++ {
			x := m.xSlopeIntercept(y)
			tview.Print(screen, "#", x+1, y+1, 1, tview.AlignCenter, tcell.ColorGreen)
		}
	}
}

func (m *MapView) handleMouseEvents(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse) {
	if action == tview.MouseLeftClick {
		// check if we clicked a City
		x, y := event.Position()
		for _, city := range cities {
			if (x >= city.X && x <= city.X+2) && (y >= city.Y && y <= city.Y+1) {
				// start travel
				m.StartTravel(city)
				//// switch to inventory view
				//gameLayout.UpdateCity(city)
				//gameLayout.ChangePage("City")
			}
		}
	}
	return action, event
}

func (m *MapView) UpdateViews() {
	// get next player location on travel path
}

func (m *MapView) StartTravel(city *City) {
	m.fromCity = gameLayout.City
	m.toCity = city

	// Calculate slope
	a := float64(m.toCity.Y-m.fromCity.Y) / float64(m.toCity.X-m.fromCity.X)
	b := float64(m.fromCity.Y) - a*float64(m.fromCity.X)
	m.ySlopeIntercept = func(x int) int {
		return int(a*float64(x) + b)
	}

	// Calculate slope
	m.xSlopeIntercept = func(y int) int {
		return int((float64(y) - b) / a)
	}
}

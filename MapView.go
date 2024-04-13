package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
)

type MapView struct {
	*tview.Box

	mapImage string

	Cities []*City
}

func NewMapView(cities []*City) *MapView {
	content, _ := os.ReadFile("drawnMap.txt")

	mapView := &MapView{
		Box:      tview.NewBox(),
		mapImage: string(content),
		Cities:   cities,
	}

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
		}
		tview.Print(screen, string(m.mapImage[i]), x, y+1, 1, tview.AlignCenter, tcell.ColorWhite)
		x++
		i++
	}

	// Draw the cities
	for _, city := range m.Cities {
		tview.Print(screen, "C", city.X, city.Y, 1, tview.AlignCenter, tcell.ColorRed)
	}
}

func handleMouseEvents(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse) {
	if action == tview.MouseLeftClick {
		x, y := event.Position()
		modal := tview.NewModal()
		modal.
			SetText(fmt.Sprintf("x: %d  y: %d", x, y)).
			AddButtons([]string{"Ok"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				app.SetRoot(grid, true).SetFocus(grid)
			})

		if err := app.SetRoot(modal, false).SetFocus(modal).Run(); err != nil {
			panic(err)
		}
	}
	return action, event
}

func handleButtonEvents(event *tcell.EventKey) *tcell.EventKey {
	//if event.Key() == tcell.KeyEnter {
	//	// save the drawnmap to a file
	//	file, _ := os.Create("drawnMap.txt")
	//	defer file.Close()
	//	for j := 0; j < 45; j++ {
	//		for i := 0; i < 160; i++ {
	//			if drawnMap[i][j] {
	//				file.WriteString("x")
	//			} else {
	//				file.WriteString(" ")
	//			}
	//		}
	//		file.WriteString("\n")
	//	}
	//}
	return event
}

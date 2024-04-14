package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"image"
	"io"
	"math"
	"slices"
)

type MapView struct {
	*tview.Box

	image string

	Cities []*City
	Player *Player

	fromCity *City
	toCity   *City

	travelPath     [](image.Point)
	playerPosition int
}

func NewMapView(cities []*City, player *Player) *MapView {
	file, _ := payload.Open("payload/drawnMap.txt")
	content, _ := io.ReadAll(file)

	mapView := &MapView{
		Box:    tview.NewBox(),
		image:  string(content),
		Cities: cities,
		Player: player,
	}

	mapView.SetMouseCapture(mapView.handleMouseEvents)

	return mapView
}

func (m *MapView) Draw(screen tcell.Screen) {
	m.Box.DrawForSubclass(screen, m)
	rectX, rectY, width, height := m.Box.GetInnerRect()

	// Scale the image
	scaledImage := scaleImage(width, height, m.image)
	// Scale the width and height
	scaleX := math.Max(1.0, 160.0/float64(width))
	scaleY := math.Max(1.0, 52.0/float64(height))

	// Draw the map
	x, y, i, imageLen := rectX, rectY, 0, len(scaledImage)
	for true {
		if i >= imageLen-1 {
			break
		}
		if scaledImage[i] == '\n' || width+rectX == x {
			x = rectX
			y++
			i++
		}
		tview.Print(screen, string(scaledImage[i]), x, y, 1, tview.AlignCenter, tcell.ColorWhite)
		x++
		i++
	}

	// Draw the Cities
	for _, city := range m.Cities {
		color := tcell.ColorRed
		if city == gameLayout.City {
			color = tcell.ColorBlue
		}

		// scale cities to map
		cX := int(float64(city.X)/scaleX) + 1
		cY := int(float64(city.Y)/scaleY) + 1

		tview.Print(screen, "\u256d", cX, cY, 1, tview.AlignCenter, color)
		tview.Print(screen, "\u2500", cX+1, cY, 1, tview.AlignCenter, color)
		tview.Print(screen, "\u256e", cX+2, cY, 1, tview.AlignCenter, color)
		tview.Print(screen, "\u2570", cX, cY+1, 1, tview.AlignCenter, color)
		tview.Print(screen, "\u2500", cX+1, cY+1, 1, tview.AlignCenter, color)
		tview.Print(screen, "\u256f", cX+2, cY+1, 1, tview.AlignCenter, color)
	}

	if m.Player.IsTravelling {
		// Draw to path to the City
		for _, p := range m.travelPath {
			tview.Print(screen, "#", int(float64(p.X)/scaleX)+2, int(float64(p.Y)/scaleY)+2, 1, tview.AlignCenter, tcell.ColorGreen)
		}

		// Draw player
		tview.Print(screen, "X", int(float64(m.travelPath[m.playerPosition].X)/scaleX)+2, int(float64(m.travelPath[m.playerPosition].Y)/scaleY)+2, 1, tview.AlignCenter, tcell.ColorRed)

		// move player
		if m.playerPosition >= len(m.travelPath)-1 {
			// Arrived at destination. Reset everything
			m.Player.IsTravelling = false
			m.playerPosition = 0
			gameLayout.UpdateCity(m.toCity)
			gameLayout.ChangePage("City")
		} else {
			m.playerPosition++
		}
	}
}

func (m *MapView) handleMouseEvents(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse) {

	_, _, width, height := m.Box.GetInnerRect()

	if action == tview.MouseLeftClick && m.Player.IsTravelling == false {
		// check if we clicked a City
		x, y := event.Position()
		for _, city := range cities {
			// scale cities to map
			scaleX := math.Max(1.0, 160.0/float64(width))
			scaleY := math.Max(1.0, 52.0/float64(height))

			cx := int(float64(city.X)/scaleX) + 1
			cy := int(float64(city.Y)/scaleY) + 1

			if (x >= cx && x <= cx+2) && (y >= cy && y <= cy+1) {
				if gameLayout.City != city {
					// start travel
					m.StartTravel(city)
				} else {
					// switch to inventory view
					gameLayout.UpdateCity(city)
					gameLayout.ChangePage("City")
				}
			}
		}
	}
	return action, event
}

func (m *MapView) StartTravel(city *City) {
	m.fromCity = gameLayout.City
	m.toCity = city

	m.calculateTravelPath()
	m.Player.IsTravelling = true
}

func (m *MapView) calculateTravelPath() {
	// Calculate slope
	a := float64(m.toCity.Y-m.fromCity.Y) / float64(m.toCity.X-m.fromCity.X)
	b := float64(m.fromCity.Y) - a*float64(m.fromCity.X)
	ySlopeIntercept := func(x int) int {
		return int(a*float64(x) + b)
	}
	xSlopeIntercept := func(y int) int {
		return int((float64(y) - b) / a)
	}

	// Draw line between cities along the x-axis
	x1 := m.fromCity.X
	x2 := m.toCity.X
	if m.fromCity.X > m.toCity.X {
		x1 = m.toCity.X
		x2 = m.fromCity.X
	}
	pathX := make([]image.Point, 0)
	for x := x1; x < x2; x++ {
		pathX = append(pathX, image.Point{x, ySlopeIntercept(x)})
	}
	// Draw line between cities along the y-axis
	y1 := m.fromCity.Y
	y2 := m.toCity.Y
	if m.fromCity.Y > m.toCity.Y {
		y1 = m.toCity.Y
		y2 = m.fromCity.Y
	}
	pathY := make([]image.Point, 0)
	for y := y1; y < y2; y++ {
		skip := false
		point := image.Point{xSlopeIntercept(y), y}
		// Check for duplicate points
		for _, pX := range pathX {
			if pX.X == point.X && pX.Y == point.Y {
				skip = true
			}
		}
		if skip {
			continue
		}

		pathY = append(pathY, point)
	}

	// Combine paths
	m.travelPath = make([]image.Point, 0)
	point := image.Point{x1, y1}
	for len(pathX) > 0 && len(pathY) > 0 {
		pointA := pathX[0]
		pointB := pathY[0]

		// Determine if pointA or pointB is closes to point
		if point.Sub(pointA).X*point.Sub(pointA).X+point.Sub(pointA).Y*point.Sub(pointA).Y < point.Sub(pointB).X*point.Sub(pointB).X+point.Sub(pointB).Y*point.Sub(pointB).Y {
			m.travelPath = append(m.travelPath, pointA)
			pathX = pathX[1:]
		} else {
			m.travelPath = append(m.travelPath, pointB)
			pathY = pathY[1:]
		}
	}
	// Append remaining points
	m.travelPath = append(m.travelPath, pathX...)
	m.travelPath = append(m.travelPath, pathY...)

	// Check direction and reverse if needed
	if m.travelPath[0].X == m.toCity.X && m.travelPath[0].Y == m.toCity.Y {
		slices.Reverse(m.travelPath)
	}
}

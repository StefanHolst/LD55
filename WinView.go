package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"io"
	"math"
	"strings"
)

type WinView struct {
	*tview.Box

	image string
}

func NewWinView() *WinView {
	file, _ := payload.Open("payload/Win.txt")
	content, _ := io.ReadAll(file)

	winView := &WinView{
		Box:   tview.NewBox(),
		image: string(content),
	}

	return winView
}

func (w *WinView) Draw(screen tcell.Screen) {
	w.Box.DrawForSubclass(screen, w)
	rectX, rectY, width, height := w.Box.GetInnerRect()

	scaledImage := scaleImage(width, height, w.image)

	// Draw the image
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
}

func scaleImage(width int, height int, image string) string {
	// find the first newline to determine the width
	actualWidth := strings.Index(image, "\n")
	actualHeight := len(image) / actualWidth //strings.Count(image, "\n")

	lines := strings.Split(image, "\n")

	// scale the image
	scaleWidth := math.Max(1, float64(actualWidth)/float64(width))
	scaleHeight := math.Max(1, float64(actualHeight)/float64(height))

	newImage := ""
	// scale the image
	for y := 0.0; y < float64(len(lines)); y += scaleHeight {
		line := lines[int(y)]
		if line == "" {
			continue
		}

		newLine := ""
		for x := 0.0; x < float64(actualWidth); x += scaleWidth {
			if len(newLine) >= width {
				break
			}
			newLine += string(line[int(x)])
		}
		newImage += newLine + "\n"
	}

	return newImage
}

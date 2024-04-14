package main

import (
	"fmt"
	"github.com/rivo/tview"
	"strings"
)

type PlayerStatsView struct {
	tview.Grid

	primitives map[string]tview.Primitive

	player *Player
}

func NewPlayerStatsView(player *Player) *PlayerStatsView {
	psv := &PlayerStatsView{
		Grid:       *tview.NewGrid(),
		primitives: make(map[string]tview.Primitive),
		player:     player,
	}

	psv.SetRows(2, 1, 1, 1)
	psv.SetColumns(-2, 0)

	psv.primitives["time"] = tview.NewTextView()
	psv.primitives["tradedBalance"] = tview.NewTextView()
	psv.primitives["account"] = tview.NewTextView()
	psv.primitives["artifacts"] = tview.NewTextView()

	// Time
	psv.AddItem(tview.NewTextView().SetText("Time"), 0, 0, 1, 1, 0, 0, false)
	psv.AddItem(psv.primitives["time"], 0, 1, 1, 1, 0, 0, false)

	psv.AddItem(tview.NewTextView().SetText("Total Traded"), 1, 0, 1, 1, 0, 0, false)
	psv.AddItem(psv.primitives["tradedBalance"], 1, 1, 1, 1, 0, 0, false)

	psv.AddItem(tview.NewTextView().SetText("Balance"), 2, 0, 1, 1, 0, 0, false)
	psv.AddItem(psv.primitives["account"], 2, 1, 1, 1, 0, 0, false)

	// List artifacts
	psv.AddItem(tview.NewTextView().SetText("Artifacts"), 4, 0, 1, 1, 0, 0, false)
	psv.AddItem(psv.primitives["artifacts"], 4, 1, 1, 1, 0, 0, false)
	return psv
}

func (psv *PlayerStatsView) UpdateViews() {
	psv.primitives["time"].(*tview.TextView).SetText(fmt.Sprintf("%s", gameLayout.time.Format("15:04")))
	psv.primitives["tradedBalance"].(*tview.TextView).SetText(fmt.Sprintf("%f", player.tradedBalance))
	psv.primitives["account"].(*tview.TextView).SetText(fmt.Sprintf("%f", player.Account))

	hasArtifacts := []string{}
	for i := 0; i < len(ArtifactTypeToString); i++ {
		if player.Artifacts[ArtifactType(i)] != nil {
			hasArtifacts = append(hasArtifacts, "[x]")
		} else {
			hasArtifacts = append(hasArtifacts, "[ ]")
		}
	}
	psv.primitives["artifacts"].(*tview.TextView).SetText(strings.Join(hasArtifacts, " "))
}

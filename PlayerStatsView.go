package main

import (
	"fmt"
	"github.com/rivo/tview"
	"math"
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
	psv.primitives["summonedBalance"] = tview.NewTextView()
	psv.primitives["suspicion"] = tview.NewTextView()
	psv.primitives["account"] = tview.NewTextView()

	// Time
	psv.AddItem(psv.primitives["time"], 0, 0, 1, 1, 0, 0, false)

	// Village suspicion bar
	psv.AddItem(tview.NewTextView().SetText("Villagers Suspicious"), 1, 0, 1, 1, 0, 0, false)
	psv.AddItem(psv.primitives["suspicion"], 1, 1, 1, 1, 0, 0, false)

	psv.AddItem(tview.NewTextView().SetText("Total Traded"), 2, 0, 1, 1, 0, 0, false)
	psv.AddItem(psv.primitives["tradedBalance"], 2, 1, 1, 1, 0, 0, false)

	psv.AddItem(tview.NewTextView().SetText("Total Summoned Traded"), 3, 0, 1, 1, 0, 0, false)
	psv.AddItem(psv.primitives["summonedBalance"], 3, 1, 1, 1, 0, 0, false)

	psv.AddItem(tview.NewTextView().SetText("Balance"), 4, 0, 1, 1, 0, 0, false)
	psv.AddItem(psv.primitives["account"], 4, 1, 1, 1, 0, 0, false)

	return psv
}

func (psv *PlayerStatsView) UpdateViews() {
	playerSuspicion := (player.summonedBalance / player.tradedBalance) * 100
	if math.IsNaN(playerSuspicion) {
		playerSuspicion = 0
	}
	psv.primitives["time"].(*tview.TextView).SetText(fmt.Sprintf("%s", gameLayout.time.Format("15:04:05")))
	psv.primitives["suspicion"].(*tview.TextView).SetText(fmt.Sprintf("%d", int(playerSuspicion)))
	psv.primitives["tradedBalance"].(*tview.TextView).SetText(fmt.Sprintf("%f", player.tradedBalance))
	psv.primitives["summonedBalance"].(*tview.TextView).SetText(fmt.Sprintf("%f", player.summonedBalance))
	psv.primitives["account"].(*tview.TextView).SetText(fmt.Sprintf("%f", player.Account))
}

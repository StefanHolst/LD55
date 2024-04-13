package main

import (
	"math"
	"math/rand"
)

type Products int

var ProductsToString = map[Products]string{
	Wood:    "Wood",
	Stone:   "Stone",
	Iron:    "Iron",
	Gold:    "Gold",
	Meat:    "Meat",
	Fish:    "Fish",
	Bread:   "Bread",
	Beer:    "Beer",
	Wine:    "Wine",
	Cloth:   "Cloth",
	Leather: "Leather",
}

const (
	Wood Products = iota
	Stone
	Iron
	Gold
	Meat
	Fish
	Bread
	Beer
	Wine
	Cloth
	Leather
)

type TradingItem struct {
	Value        float64
	Amount       int
	ScalingValue int
	IsSummoned   bool
}

func NewTradingItem() *TradingItem {
	return &TradingItem{
		ScalingValue: rand.Intn(2000) + 200,
	}
}

func (t *TradingItem) Add(amount int) {
	t.Amount += amount
	t.ScaleValue()
}

func (t *TradingItem) Remove(amount int) {
	if (t.Amount - amount) < 0 {
		t.Amount = 0
	} else {
		t.Amount -= amount
	}
	t.ScaleValue()
}

func (t *TradingItem) ScaleValue() {
	t.Value = math.Pow(float64(t.Amount+1), -0.3) * float64(t.ScalingValue)
}

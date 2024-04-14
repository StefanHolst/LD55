package main

import "math/rand"

type Player struct {
	// inventory
	Inventory map[ProductType]*Product

	Account  float64
	Name     string
	HomeCity string

	tradedBalance   float64
	summonedBalance float64

	positionX, positionY int
}

func NewPlayer(name string) *Player {
	player := &Player{
		Inventory: NewProducts(),
		Name:      name,
		Account:   1000,
	}

	return player
}

func (p *Player) Buy(city *City, productType ProductType) {
	// check if the Player has enough money
	if p.Account > city.Inventory[productType].Value && len(city.Inventory[productType].Items) > 0 {
		// Deduct the money from the Player account
		p.Account -= city.Inventory[productType].Value

		// Add to the traded sums
		p.tradedBalance += city.Inventory[productType].Value

		// Pop the item from the City inventory
		item := city.Inventory[productType].PopItem()

		// Add the item to the Player inventory
		p.Inventory[productType].AddItem(item)
	}
}

func (p *Player) Sell(city *City, productType ProductType) {
	// check if the Player has enough productType
	if len(p.Inventory[productType].Items) > 0 {
		// Add the money to the Player account
		p.Account += city.Inventory[productType].Value * 0.8

		// Add to the traded sums
		p.tradedBalance += city.Inventory[productType].Value

		// Pop the item from the Player inventory
		item := p.Inventory[productType].PopItem()

		// Add the item to the City inventory
		city.Inventory[productType].AddItem(item)

		// if the item is summoned, add to the summoned balance
		if item.IsSummoned {
			p.summonedBalance += city.Inventory[productType].Value
		}
	}
}

func (p *Player) SummonProducts() {
	for i := 0; i < 10; i++ {
		for j := 0; j < rand.Intn(1000); j++ {
			p.Inventory[ProductType(i)].AddItem(NewTradingItem(true, p.Inventory[ProductType(i)]))
		}
	}
}

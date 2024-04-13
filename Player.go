package main

type Player struct {
	// inventory
	Inventory map[Products]*TradingItem

	Account  float64
	Name     string
	HomeCity string

	// Summoned items?
}

func NewPlayer(name string) *Player {
	player := &Player{
		Inventory: make(map[Products]*TradingItem),
		Name:      name,
		Account:   1000,
	}

	for i := 0; i < 10; i++ {
		player.Inventory[Products(i)] = NewTradingItem()
	}
	return player
}

func (p *Player) buy(city *City, products Products) {
	// check if the player has enough money
	if p.Account > city.TradingItems[products].Value && city.TradingItems[products].Amount > 0 {
		p.Inventory[products].Amount++
		p.Account -= city.TradingItems[products].Value
		city.TradingItems[products].Remove(1)
	}
}

func (p *Player) sell(city *City, products Products) {
	// check if the player has enough products
	if p.Inventory[products].Amount > 0 {
		p.Inventory[products].Remove(1)
		p.Account += city.TradingItems[products].Value
		city.TradingItems[products].Add(1)
	}
}

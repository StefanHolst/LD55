package main

import "math/rand"

type City struct {
	TradingItems map[Products]*TradingItem

	Usage      map[Products]int
	Production map[Products]int

	Name       string
	Population int

	X, Y int
}

func NewCity(name string, population int, x int, y int) *City {
	city := &City{
		TradingItems: make(map[Products]*TradingItem),
		Usage:        make(map[Products]int),
		Production:   make(map[Products]int),
		Name:         name,
		Population:   population,
		X:            x,
		Y:            y,
	}

	for i := 0; i < 10; i++ {
		city.TradingItems[Products(i)] = NewTradingItem()
		city.TradingItems[Products(i)].Amount = rand.Intn(100)
		city.Production[Products(i)] = rand.Intn(5)
		city.Usage[Products(i)] = rand.Intn(5)
	}

	return city
}

func (c *City) UpdateProduction() {
	for i := 0; i < 10; i++ {
		c.TradingItems[Products(i)].Add(c.Production[Products(i)])
	}
}

func (c *City) UpdateUsage() {
	for i := 0; i < 10; i++ {
		c.TradingItems[Products(i)].Remove(c.Usage[Products(i)])
	}
}

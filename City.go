package main

import "math/rand"

type City struct {
	Inventory map[ProductType]*Product

	Usage      map[ProductType]int
	Production map[ProductType]int

	Name       string
	Population int

	X, Y int
}

func NewCity(name string, population int, x int, y int) *City {
	city := &City{
		Inventory:  NewProducts(),
		Usage:      make(map[ProductType]int),
		Production: make(map[ProductType]int),
		Name:       name,
		Population: population,
		X:          x,
		Y:          y,
	}

	for i := 0; i < 10; i++ {
		city.Inventory[ProductType(i)].Items = make([]*TradingItem, rand.Intn(100))
		city.Production[ProductType(i)] = rand.Intn(5)
		city.Usage[ProductType(i)] = rand.Intn(5)
	}

	return city
}

func (c *City) UpdateProduction() {
	for i := 0; i < 10; i++ {
		produced := c.Production[ProductType(i)]
		for j := 0; j < produced; j++ {
			c.Inventory[ProductType(i)].AddItem(NewTradingItem(false, c.Inventory[ProductType(i)]))
		}
	}
}

func (c *City) UpdateUsage() {
	for i := 0; i < 10; i++ {
		used := c.Usage[ProductType(i)]
		for j := 0; j < used; j++ {
			c.Inventory[ProductType(i)].PopItem()
		}
	}
}

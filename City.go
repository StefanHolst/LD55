package main

import "math/rand"

type City struct {
	Inventory map[ProductType]*Product

	BaselineUsage      map[ProductType]int
	Usage              map[ProductType]int
	BaselineProduction map[ProductType]int
	Production         map[ProductType]int

	Artifact     *Artifact
	ArtifactSold bool

	Name       string
	Population int

	X, Y int
}

func NewCity(name string, population int, x int, y int, artifact *Artifact) *City {
	city := &City{
		Inventory:          NewProducts(),
		BaselineUsage:      make(map[ProductType]int),
		Usage:              make(map[ProductType]int),
		BaselineProduction: make(map[ProductType]int),
		Production:         make(map[ProductType]int),
		Name:               name,
		Population:         population,
		X:                  x,
		Y:                  y,
		Artifact:           artifact,
	}

	for i := 0; i < len(ProductTypeToString); i++ {
		city.Inventory[ProductType(i)].Amount = rand.Intn(100)
		city.BaselineProduction[ProductType(i)] = rand.Intn(5)
		city.Production[ProductType(i)] = rand.Intn(5)
		city.BaselineUsage[ProductType(i)] = rand.Intn(5)
		city.Usage[ProductType(i)] = rand.Intn(5)
	}

	return city
}

func (c *City) UpdateProduction() {
	for i := 0; i < len(ProductTypeToString); i++ {
		produced := c.Production[ProductType(i)] + c.BaselineProduction[ProductType(i)]
		for j := 0; j < produced; j++ {
			c.Inventory[ProductType(i)].Add()
		}
	}
}

func (c *City) UpdateUsage() {
	for i := 0; i < len(ProductTypeToString); i++ {
		used := c.Usage[ProductType(i)] + c.BaselineUsage[ProductType(i)]
		for j := 0; j < used; j++ {
			c.Inventory[ProductType(i)].Remove()
		}
	}
}

func (c *City) UpdateProductionScale() {
	for i := 0; i < len(ProductTypeToString); i++ {
		c.Production[ProductType(i)] = rand.Intn(5)
	}
}

func (c *City) UpdateUsageScale() {
	for i := 0; i < len(ProductTypeToString); i++ {
		c.Usage[ProductType(i)] = rand.Intn(5)
	}
}

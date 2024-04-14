package main

type Player struct {
	// inventory
	Inventory map[ProductType]*Product

	Artifacts map[ArtifactType]*Artifact

	Account float64
	Name    string

	tradedBalance float64

	IsTravelling bool
}

func NewPlayer(name string) *Player {
	player := &Player{
		Inventory: NewProducts(),
		Artifacts: map[ArtifactType]*Artifact{},
		Name:      name,
		Account:   1000,
	}

	return player
}

func (p *Player) Buy(city *City, productType ProductType) {
	// check if the Player has enough money
	if p.Account > city.Inventory[productType].Value && city.Inventory[productType].Amount > 0 {
		// Deduct the money from the Player account
		p.Account -= city.Inventory[productType].Value

		// Add to the traded sums
		p.tradedBalance += city.Inventory[productType].Value

		// Pop the item from the City inventory
		city.Inventory[productType].Remove()

		// Add the item to the Player inventory
		p.Inventory[productType].Add()
	}
}

func (p *Player) Sell(city *City, productType ProductType) {
	// check if the Player has enough productType
	if p.Inventory[productType].Amount > 0 {
		// Add the money to the Player account
		p.Account += city.Inventory[productType].Value * 0.8

		// Add to the traded sums
		p.tradedBalance += city.Inventory[productType].Value

		// Pop the item from the Player inventory
		p.Inventory[productType].Remove()

		// Add the item to the City inventory
		city.Inventory[productType].Add()
	}
}

func (p *Player) BuyArtifact(city *City) bool {
	if city.ArtifactSold || p.Account < city.Artifact.Value {
		return false
	}

	// Deduct the money from the Player account
	p.Account -= city.Artifact.Value

	// Add the artifact to the Player inventory
	p.Artifacts[city.Artifact.ArtifactType] = city.Artifact

	// Set the artifact as sold
	city.ArtifactSold = true

	// Check if all artifacts are sold
	for _, city := range gameLayout.Cities {
		if !city.ArtifactSold {
			return true
		}
	}

	showWin()
	return true
}

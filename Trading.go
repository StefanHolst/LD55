package main

import (
	"math"
	"math/rand"
)

var ProductTypeToString = map[ProductType]string{
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
	Wood ProductType = iota
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

type ProductType int

var ArtifactTypeToString = map[ArtifactType]string{
	TheHornOfFreyja:       "The Horn of Freyja",
	TheSakuraKimono:       "The Sakura Kimono",
	TheStatuesCompass:     "The Statues Compass",
	TheLuckyChip:          "The Lucky Chip",
	TheCarnivalMask:       "The Carnival Mask",
	TheDiamondKey:         "The Diamond Key",
	TheWinterCloak:        "The Winter Cloak",
	TheDreamtimeBoomerang: "The Dreamtime Boomerang",
}

const (
	TheHornOfFreyja ArtifactType = iota
	TheSakuraKimono
	TheStatuesCompass
	TheLuckyChip
	TheCarnivalMask
	TheDiamondKey
	TheWinterCloak
	TheDreamtimeBoomerang
)

type ArtifactType int

type Artifact struct {
	ArtifactType ArtifactType
	Value        float64
}

func NewArtifact(artifactType ArtifactType) *Artifact {
	return &Artifact{
		ArtifactType: artifactType,
		Value:        (rand.Float64() + 1) * 10000,
	}
}

type Product struct {
	ProductType  ProductType
	Value        float64
	ScalingValue int
	Amount       int
	//Weight       float64
	//Items []*TradingItem
}

func NewProduct(productType ProductType) *Product {
	return &Product{
		ProductType:  productType,
		ScalingValue: rand.Intn(2000) + 200,
	}
}

func NewProducts() map[ProductType]*Product {
	products := make(map[ProductType]*Product)
	for i := 0; i < len(ProductTypeToString); i++ {
		products[ProductType(i)] = NewProduct(ProductType(i))
	}
	return products
}

func (p *Product) updateScale() {
	p.Value = math.Pow(float64(p.Amount+1), -0.3) * float64(p.ScalingValue)
}

func (p *Product) Add() {
	p.Amount++
	p.updateScale()
}

func (p *Product) Remove() {
	if p.Amount == 0 {
		return
	}
	p.Amount--
	p.updateScale()
}

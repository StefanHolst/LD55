package main

import (
	"math"
	"math/rand"
)

type ProductType int

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

type Product struct {
	ProductType  ProductType
	Value        float64
	ScalingValue int
	//Weight       float64
	Items []*TradingItem
}

func NewProduct(productType ProductType) *Product {
	return &Product{
		ProductType:  productType,
		ScalingValue: rand.Intn(2000) + 200,
	}
}

func NewProducts() map[ProductType]*Product {
	products := make(map[ProductType]*Product)
	for i := 0; i < 10; i++ {
		products[ProductType(i)] = NewProduct(ProductType(i))
	}
	return products
}

func (p *Product) updateScale() {
	p.Value = math.Pow(float64(len(p.Items)+1), -0.3) * float64(p.ScalingValue)
}

func (p *Product) AddItem(item *TradingItem) {
	p.Items = append(p.Items, item)
	p.updateScale()
}

func (p *Product) PopItem() *TradingItem {
	if len(p.Items) == 0 {
		return nil
	}
	item := p.Items[0]
	p.Items = p.Items[1:]
	p.updateScale()
	return item
}

type TradingItem struct {
	IsSummoned bool
	Product    *Product
}

func NewTradingItem(isSummoned bool, product *Product) *TradingItem {
	return &TradingItem{
		IsSummoned: isSummoned,
		Product:    product,
	}
}

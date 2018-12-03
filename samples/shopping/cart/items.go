package cart

import "time"

type Item struct {
	Name           string
	Units          int
	PricePerUnit   float32
	ExpirationDate time.Time
}

var Items = []*Item{
	{
		Name:           "Lettuce",
		Units:          2,
		PricePerUnit:   0.97,
		ExpirationDate: time.Now().Add(-72 * time.Hour),
	},
	{
		Name:           "Onions",
		Units:          5,
		PricePerUnit:   0.27,
		ExpirationDate: time.Now().Add(72 * time.Hour),
	},
	{
		Name:           "Milk",
		Units:          4,
		PricePerUnit:   1.23,
		ExpirationDate: time.Now().Add(48 * time.Hour),
	},
	{
		Name:           "Eggs Free range",
		Units:          2,
		PricePerUnit:   2.25,
		ExpirationDate: time.Now().Add(96 * time.Hour),
	},
	{
		Name:           "Tomatoes",
		Units:          6,
		PricePerUnit:   0.45,
		ExpirationDate: time.Now().Add(-1 * time.Hour),
	},
	{
		Name:           "Avocado",
		Units:          7,
		PricePerUnit:   1.05,
		ExpirationDate: time.Now().Add(1 * time.Hour),
	},
}

var Items2 = []*Item{
	{
		Name:           "Lettuce",
		Units:          4,
		PricePerUnit:   0.97,
		ExpirationDate: time.Now().Add(-2 * time.Hour),
	},
	{
		Name:           "Potatoes",
		Units:          7,
		PricePerUnit:   0.45,
		ExpirationDate: time.Now().Add(72 * time.Hour),
	},
	{
		Name:           "Kale",
		Units:          1,
		PricePerUnit:   1.23,
		ExpirationDate: time.Now().Add(48 * time.Hour),
	},
	{
		Name:           "Cucumber",
		Units:          3,
		PricePerUnit:   0.25,
		ExpirationDate: time.Now().Add(96 * time.Hour),
	},
	{
		Name:           "Carrots",
		Units:          9,
		PricePerUnit:   0.17,
		ExpirationDate: time.Now().Add(1 * time.Hour),
	},
	{
		Name:           "Tofu",
		Units:          1,
		PricePerUnit:   3.05,
		ExpirationDate: time.Now().Add(1 * time.Hour),
	},
}

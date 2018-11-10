package structure

type itemType int

const (
	bakery itemType = iota
	fruits
	vegetables
	drinks
	frozen
	healthAndBeauty
)

type shoppingItem struct {
	Name     string
	Price    float32
	ItemType itemType
}

func newShoppingItem(name string, itemType itemType, price float32) *shoppingItem {
	return &shoppingItem{
		Name:     name,
		ItemType: itemType,
		Price:    price,
	}
}

func currentShoppingCart() []*shoppingItem {
	return []*shoppingItem{
		newShoppingItem("Bread", bakery, 0.78),
		newShoppingItem("Watermelon", fruits, 1.23),
		newShoppingItem("Deodorant", healthAndBeauty, 2.25),
		newShoppingItem("Pack 6 Cokes", drinks, 4.10),
		newShoppingItem("Red wine", drinks, 7.10),
		newShoppingItem("Vegetarian Pizza", frozen, 3.18),
	}
}

package structure

type ItemType int

const (
	Bakery ItemType = iota
	Fruits
	Vegetables
	Drinks
	Frozen
	HealthAndBeauty
	FrozenFood
)

type ShoppingItem struct {
	Name     string
	Price    float32
	ItemType ItemType
}

func NewShoppingItem(name string, itemType ItemType, price float32) *ShoppingItem {
	return &ShoppingItem{
		Name:     name,
		ItemType: itemType,
		Price:    price,
	}
}

func currentShoppingCart() []*ShoppingItem {
	return []*ShoppingItem{
		NewShoppingItem("Bread", Bakery, 0.78),
		NewShoppingItem("Watermelon", Fruits, 1.23),
		NewShoppingItem("Deodorant", HealthAndBeauty, 2.25),
		NewShoppingItem("Pack 6 Cokes", Drinks, 4.10),
		NewShoppingItem("Red wine", Drinks, 7.10),
		NewShoppingItem("Vegetarian Pizza", Frozen, 3.18),
	}
}

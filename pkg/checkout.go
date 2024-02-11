package pkg

type ICheckout interface {
	Scan(item string)
	GetTotalPrice() int
}

type Checkout struct {
	basket map[string]*int
}

// NewCheckout creates instance of Checkout
func NewCheckout() *Checkout {
	return &Checkout{
		basket: map[string]*int{},
	}
}

// Scan adds item to map of item and quantity. If item already exists in the map, then the quantity
// is incremented. Since we are storing the quantity reference in the map, we can just increment the value without
// storing it again in the map.
func (c *Checkout) Scan(item string) {
	if q, ok := c.basket[item]; ok {
		*q++
	} else {
		quantity := 1
		if _, ok := Items[item]; ok {
			c.basket[item] = &quantity
		}

	}
}

func (c *Checkout) GetTotalPrice() int {
	// TODO implement me
	return -1
}

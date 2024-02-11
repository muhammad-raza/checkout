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

// GetTotalPrice loops through basket and add item price to the total value. If offer exists for the particular item,
// then add the discounted price. e.g. if `A` cost 50, and we add 7 'A's to the basket, and there is an offer that
// 3 'A's will cost 130. We will add 120 to the total and deduce 3 from 7 quantity in the basket. Keep repeating this
// process until the remaining quantity is less than the offer quantity. In this example, remainder will be 1 which is
// less than offer. Then add full price of 1 'A' to the total. The total after offer should be 310.
func (c *Checkout) GetTotalPrice() int {
	total := 0
	for k, v := range c.basket {
		quantity := *v
		desc := Items[k]
		if desc.Offer != nil {
			for quantity >= desc.Offer.Quantity {
				total = total + desc.Offer.Price
				quantity = quantity - desc.Offer.Quantity
			}
		}
		total = total + (desc.UnitPrice * quantity)
	}
	return total
}

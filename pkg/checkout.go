package pkg

type ICheckout interface {
	Scan(item string)
	GetTotalPrice() int
}

type Checkout struct{}

func NewCheckout() *Checkout {
	return &Checkout{}
}

func (c *Checkout) Scan(item string) {
	// TODO implement me
}

func (c *Checkout) GetTotalPrice() int {
	// TODO implement me
	return -1
}

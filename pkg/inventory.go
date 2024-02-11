package pkg

type ItemDescription struct {
	UnitPrice int
	Offer     *Offer
}

type Offer struct {
	Quantity int
	Price    int
}

var Items = map[string]ItemDescription{
	"A": {
		UnitPrice: 50,
		Offer: &Offer{
			Quantity: 3,
			Price:    130,
		},
	},
	"B": {
		UnitPrice: 30,
		Offer: &Offer{
			Quantity: 2,
			Price:    45,
		},
	},
	"C": {
		UnitPrice: 20,
	},
	"D": {
		UnitPrice: 15,
	},
}

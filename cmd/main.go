package main

import (
	"fmt"

	"github.com/manifoldco/promptui"

	"checkout/pkg"
)

type item struct {
	Name      string
	UnitPrice int
	Offer     string
}

func main() {
	order := []string{"A", "B", "C", "D"}
	items := []item{}
	for _, o := range order {
		desc := pkg.Items[o]
		it := item{
			Name:      o,
			UnitPrice: desc.UnitPrice,
		}
		if desc.Offer != nil {
			it.Offer = fmt.Sprintf("%d for price %d", desc.Offer.Quantity, desc.Offer.Price)
		} else {
			it.Offer = "No offer"
		}
		items = append(items, it)
	}
	items = append(items, item{
		Name: "Exit",
	})

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F34C {{ .Name | cyan }} ({{ .UnitPrice | red }})",
		Inactive: "{{ .Name | cyan }} ({{ .UnitPrice | red }})",
		Selected: "\U0001F34C {{ .Name | red | cyan }}",
		Details: `
--------- Item details ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Unit Price:" | faint }}	{{ .UnitPrice }}
{{ "Offer:" | faint }}	{{ .Offer }}`,
	}

	prompt := promptui.Select{
		Label:     "Add to basket",
		Items:     items,
		Templates: templates,
	}

	checkout := pkg.NewCheckout()

	i := -1
	for i < len(items)-1 {
		if i > -1 {
			name := order[i]
			checkout.Scan(name)
		}
		printCheckout(checkout)
		i, _, _ = prompt.Run()
	}
	printCheckout(checkout)
}

func printCheckout(checkout *pkg.Checkout) {
	basket := checkout.GetBasket()
	for k, v := range basket {
		fmt.Printf("Item: %s, Quantity: %d\n", k, *v)
	}
	fmt.Printf("Basket total: %d\n\n", checkout.GetTotalPrice())
}

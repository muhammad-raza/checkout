## Checkout

### Improvements

- Currently, the inventory data is in `inventory.go` which can be moved to database
- Item only supports one offer i.e. `3 for 130`. There is room to improve it and add more offer types e.g. `Meal Deal`
  or `3 for the price of 2` e.t.c
- Unit price is detached from offer price. e.g. if the unit price is changed, the offer can become invalid. This can be
  linked by introducing percentage reduction in price.
- To avoid calculating total price again and again. we can add `totalPrice` and a `priceChanged` attributes
  to `Checkout` struct. We can then return `totalPrice` unless `priceChanged` is true. If `priceChanged` is true, then
  calculate price again. `priceChanged` flag can be set in `Scan` function.
- Nice to have a cli for visual purpose.

### Running test

To run tests, use following command

```shell
make tests
```

package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckout(t *testing.T) {
	tests := []struct {
		name          string
		items         []string
		expectedTotal int
	}{
		{
			name:          "not item added",
			items:         []string{},
			expectedTotal: 0,
		},
		{
			name:          "add single item",
			items:         []string{"A"},
			expectedTotal: 50,
		},
		{
			name:          "duplicate item should not apply offer if less than offer Quantity",
			items:         []string{"A", "A"},
			expectedTotal: 100,
		},
		{
			name:          "duplicate item should apply offer if match the offer Quantity",
			items:         []string{"A", "A", "A"},
			expectedTotal: 130,
		},
		{
			name:          "apply offer as many times as offer can be applied",
			items:         []string{"A", "A", "A", "A", "A", "A", "A"},
			expectedTotal: 310,
		},
		{
			name:          "add items at random order should apply offer",
			items:         []string{"A", "B", "A", "C", "A", "D", "A"},
			expectedTotal: 245,
		},
		{
			name:          "should not apply offers to non offer items not matter how many times its added",
			items:         []string{"C", "C", "C", "C", "C", "C", "C"},
			expectedTotal: 140,
		},
		{
			name:          "item does not exist in inventory",
			items:         []string{"E"},
			expectedTotal: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewCheckout()
			for _, item := range test.items {
				c.Scan(item)
			}
			assert.Equal(t, test.expectedTotal, c.GetTotalPrice())
		})
	}
}

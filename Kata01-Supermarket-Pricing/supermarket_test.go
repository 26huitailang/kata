package supermarket

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscountPlan(t *testing.T) {
	dp := DiscountPlan{}
	t.Run("Buy N get M free", func(t *testing.T) {
		cases := []struct {
			name     string
			expected int
			total    int
			buy      int
			free     int
		}{
			{name: "buy 2 get 1 free", expected: 67, total: 100, buy: 2, free: 1},
			{name: "buy 1 get 1 free", expected: 51, total: 101, buy: 1, free: 1},
			{name: "buy 6 get 6 free", expected: 52, total: 100, buy: 6, free: 6},
		}
		for _, tt := range cases {
			t.Run(tt.name, func(t *testing.T) {
				actual := dp.BuyNGetMFree(tt.total, tt.buy, tt.free)
				assert.Equal(t, tt.expected, actual)
			})
		}
	})

	t.Run("Buy N for M dollars", func(t *testing.T) {
		cases := []struct {
			name     string
			expected []int
			total    int
			buy      int
			free     int
		}{
			{name: "buy 3 for m dollar", expected: []int{2, 0}, total: 6, buy: 3},
			{name: "buy 3 for m dollar", expected: []int{2, 1}, total: 7, buy: 3},
			{name: "buy 3 for m dollar", expected: []int{1, 2}, total: 5, buy: 3},
		}
		for _, tt := range cases {
			t.Run(tt.name, func(t *testing.T) {
				gotDiscount, gotFullprice := dp.BuyNForMDollars(tt.total, tt.buy)
				assert.Equal(t, tt.expected[0], gotDiscount)
				assert.Equal(t, tt.expected[1], gotFullprice)
			})
		}
	})
}

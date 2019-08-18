package supermarket

type DiscountPlan struct {
}

// BuyNGetMFree return stock shows user should pay count / or stock to display with total
// if a shelf of 100 cans is priced using “buy two, get one free”, how do you value the stock?
func (dp *DiscountPlan) BuyNGetMFree(total, buyN, freeM int) (stock int) {
	remainder := total % (buyN + freeM)
	result := total / (buyN + freeM) // round to zero
	return result*buyN + remainder
}

// BuyNForMDollars return user buy N and how many pay for m dollar and how mang pay for full price
func (dp *DiscountPlan) BuyNForMDollars(total, buyN int) (discount, fullprice int) {
	fullprice = total % buyN
	discount = (total - fullprice) / buyN
	return discount, fullprice
}

type IGoods interface {
	Pricing()
}

type Supermarket struct {
}

func (super *Supermarket) Pricing() {

}

type GoodItem struct {
	ID            int
	Name          string
	PricePerPound float32
	Stock         int
}

type Goods struct {
}

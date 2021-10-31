package main

import (
	"fmt"

	"github.com/artnoi43/bond/lib"
)

func main() {
	aBond := lib.AnnualBond{lib.Bond{Par: 1000, Coupon: 10, Maturity: 4}}
	sBond := lib.SemiBond{lib.Bond{Par: 1000, Coupon: 10, Maturity: 4}}

	var rf float64 = 0.1
	fmt.Println(aBond.Price(rf))
	fmt.Println(sBond.Price(rf))
}

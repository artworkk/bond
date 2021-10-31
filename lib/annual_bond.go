package lib

type AnnualBond struct {
	Bond
}

func (b *AnnualBond) Price(rf float64) float64 {
	var price float64
	pmt := (b.Coupon / 100) * b.Par
	for i := 1; i <= b.Maturity; i++ {
		PVCoupon := PV(pmt, rf, i)
		if i == b.Maturity {
			price += PVCoupon + PV(b.Par, rf, i)
		} else {
			price += PVCoupon
		}
	}
	return price
}

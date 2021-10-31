package lib

type SemiBond struct {
	Bond
}

func (b *SemiBond) Price(rf float64) float64 {
	var price float64
	semiRf := rf / 2
	pmt := (b.Coupon / 100 / 2) * b.Par
	for i := 1; i <= b.Maturity*2; i++ {
		PVCoupon := PV(pmt, semiRf, i)
		if i == b.Maturity*2 {
			price += PVCoupon + PV(b.Par, semiRf, i)
		} else {
			price += PVCoupon
		}
	}
	return price
}

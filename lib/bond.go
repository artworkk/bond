package lib

type Bond struct {
	Par      float64 `json:"par"`
	Coupon   float64 `json:"coupon"`
	Maturity int     `json:"maturity"`
}

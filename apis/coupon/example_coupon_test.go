package coupon_test

func ExampleCouponGetCoupon() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := coupon.CouponGetCoupon(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleCouponPushCoupon() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := coupon.CouponPushCoupon(ctx, payload)
	fmt.Println(resp, err)
}

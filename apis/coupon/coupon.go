// Package coupon 优惠券接口
package coupon

const (
	apiCouponGetCoupon  = "/product/coupon/get_list"
	apiCouponPushCoupon = "/product/coupon/push"
)

/*
获取优惠券列表

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/coupon/get_coupon.html
POST https://api.weixin.qq.com/product/coupon/get_list?access_token=xxxxxxxxx
*/
func CouponGetCoupon(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCouponGetCoupon, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
发放优惠券

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/coupon/push_coupon.html
POST https://api.weixin.qq.com/product/coupon/push?access_token=xxxxxxxxx
*/
func CouponPushCoupon(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCouponPushCoupon, bytes.NewReader(payload), "application/json;charset=utf-8")
}

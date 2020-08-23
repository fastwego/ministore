// Package order 订单接口
package order

const (
	apiOrderGetOrderList   = "/product/order/get_list"
	apiOrderGetOrderDetail = "/product/order/get"
)

/*
获取订单列表

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/order/get_order_list.html
POST https://api.weixin.qq.com/product/order/get_list?access_token=xxxxxxxxx
*/
func OrderGetOrderList(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiOrderGetOrderList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取订单详情

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/order/get_order_detail.html
POST https://api.weixin.qq.com/product/order/get?access_token=xxxxxxxxx
*/
func OrderGetOrderDetail(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiOrderGetOrderDetail, bytes.NewReader(payload), "application/json;charset=utf-8")
}

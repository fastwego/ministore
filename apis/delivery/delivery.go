// Package delivery 物流接口
package delivery

const (
	apiDeliveryGetDeliveryCompanyList = "/product/delivery/get_company_list"
	apiDeliverySendDelivery           = "/product/delivery/send"
)

/*
获取快递公司列表

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/delivery/get_delivery_company_list.html
POST https://api.weixin.qq.com/product/delivery/get_company_list?access_token=xxxxxxxxx
*/
func DeliveryGetDeliveryCompanyList(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDeliveryGetDeliveryCompanyList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
订单发货

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/delivery/send_delivery.html
POST https://api.weixin.qq.com/product/delivery/send?access_token=xxxxxxxxx
*/
func DeliverySendDelivery(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDeliverySendDelivery, bytes.NewReader(payload), "application/json;charset=utf-8")
}

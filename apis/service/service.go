// Package service 服务商接口
package service

const (
	apiServiceCheckAuth           = "/product/service/check_auth"
	apiServiceGetServiceList      = "/product/service/get_list"
	apiServiceGetServiceOrderList = "/product/service/get_order_list"
)

/*
登录验证

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/service/check_auth.html
POST https://api.weixin.qq.com/product/service/check_auth?component_access_token=xxxxxxxxx
*/
func ServiceCheckAuth(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiServiceCheckAuth, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户购买的在有效期内的服务列表

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/service/get_service_list.html
POST https://api.weixin.qq.com/product/service/get_list?access_token=xxxxxxxxx
*/
func ServiceGetServiceList(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiServiceGetServiceList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户购买的服务列表

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/service/get_service_order_list.html
POST https://api.weixin.qq.com/product/service/get_order_list?access_token=xxxxxxxxx
*/
func ServiceGetServiceOrderList(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiServiceGetServiceOrderList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

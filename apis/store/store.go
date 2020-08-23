// Package store 店铺接口
package store

const (
	apiStoreGetStoreInfo = "/product/store/get_info"
)

/*
获取基本信息

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/store/get_store_info.html
POST https://api.weixin.qq.com/product/store/get_info?access_token=xxxxxxxxx
*/
func StoreGetStoreInfo(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiStoreGetStoreInfo, bytes.NewReader(payload), "application/json;charset=utf-8")
}

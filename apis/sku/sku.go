// Package sku SKU接口(修改需要重新上架商品)
package sku

const (
	apiSkuAddSku      = "/product/sku/add"
	apiSkuBatchAddSku = "/product/sku/batch_add"
	apiSkuDelSku      = "/product/sku/del"
	apiSkuGetSku      = "/product/sku/get"
	apiSkuUpSku       = "/product/sku/update"
	apiSkuUpSkuPrice  = "/product/sku/update_price"
	apiSkuUpStock     = "/product/sku/stock/update"
	apiSkuGetStock    = "/product/sku/stock/get"
)

/*
添加SKU

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/sku/add_sku.html
POST https://api.weixin.qq.com/product/sku/add?access_token=xxxxxxxxx
*/
func SkuAddSku(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSkuAddSku, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量添加SKU

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/sku/batch_add_sku.html
POST https://api.weixin.qq.com/product/sku/batch_add?access_token=xxxxxxxxx
*/
func SkuBatchAddSku(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSkuBatchAddSku, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除SKU

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/sku/del_sku.html
POST https://api.weixin.qq.com/product/sku/del?access_token=xxxxxxxxx
*/
func SkuDelSku(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSkuDelSku, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取SKU信息

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/sku/get_sku.html
POST https://api.weixin.qq.com/product/sku/get?access_token=xxxxxxxxx
*/
func SkuGetSku(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSkuGetSku, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新SKU

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/sku/up_sku.html
POST https://api.weixin.qq.com/product/sku/update?access_token=xxxxxxxxx
*/
func SkuUpSku(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSkuUpSku, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新SKU价格

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/sku/up_sku_price.html
POST https://api.weixin.qq.com/product/sku/update_price?access_token=xxxxxxxxx
*/
func SkuUpSkuPrice(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSkuUpSkuPrice, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新库存

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/sku/up_stock.html
POST https://api.weixin.qq.com/product/sku/stock/update?access_token=xxxxxxxxx
*/
func SkuUpStock(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSkuUpStock, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取库存

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/sku/get_stock.html
POST https://api.weixin.qq.com/product/sku/stock/get?access_token=xxxxxxxxx
*/
func SkuGetStock(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSkuGetStock, bytes.NewReader(payload), "application/json;charset=utf-8")
}

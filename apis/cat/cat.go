// Package cat 接入商品前必需接口
package cat

import "bytes"

const (
	apiCatGetCatList         = "/product/category/get"
	apiCatGetBrand           = "/product/brand/get"
	apiCatGetFreightTemplate = "/product/delivery/get_freight_template"
)

/*
获取类目详情

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/cat/get_cat_list.html
POST https://api.weixin.qq.com/product/category/get?access_token=xxxxxxxxx
*/
func CatGetCatList(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCatGetCatList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取品牌列表

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/cat/get_brand.html
POST https://api.weixin.qq.com/product/brand/get?access_token=xxxxxxxxx
*/
func CatGetBrand(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCatGetBrand, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取运费模板

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/cat/get_freight_template.html
POST https://api.weixin.qq.com/product/delivery/get_freight_template?access_token=xxxxxxxxx
*/
func CatGetFreightTemplate(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCatGetFreightTemplate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

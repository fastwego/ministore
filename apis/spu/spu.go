// Package spu SPU接口(修改需要重新上架商品)
package spu

const (
	apiSpuAddSpu         = "/product/spu/add"
	apiSpuDelSpu         = "/product/spu/del"
	apiSpuGetSpu         = "/product/spu/get"
	apiSpuGetSpuList     = "/product/spu/get_list"
	apiSpuUpSpu          = "/product/spu/update"
	apiSpuUpSpuListing   = "/product/spu/listing"
	apiSpuUpSpuDelisting = "/product/spu/delisting"
)

/*
添加商品

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/add_spu.html
POST https://api.weixin.qq.com/product/spu/add?access_token=xxxxxxxxx
*/
func SpuAddSpu(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSpuAddSpu, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除商品

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/del_spu.html
POST https://api.weixin.qq.com/product/spu/del?access_token=xxxxxxxxx
*/
func SpuDelSpu(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSpuDelSpu, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取商品

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/get_spu.html
POST https://api.weixin.qq.com/product/spu/get?access_token=xxxxxxxxx
*/
func SpuGetSpu(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSpuGetSpu, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取商品列表

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/get_spu_list.html
POST https://api.weixin.qq.com/product/spu/get_list?access_token=xxxxxxxxx
*/
func SpuGetSpuList(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSpuGetSpuList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新商品

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/up_spu.html
POST https://api.weixin.qq.com/product/spu/update?access_token=xxxxxxxxx
*/
func SpuUpSpu(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSpuUpSpu, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
上架商品

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/up_spu_listing.html
POST https://api.weixin.qq.com/product/spu/listing?access_token=xxxxxxxxx
*/
func SpuUpSpuListing(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSpuUpSpuListing, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
下架商品

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/up_spu_delisting.html
POST https://api.weixin.qq.com/product/spu/delisting?access_token=xxxxxxxxx
*/
func SpuUpSpuDelisting(ctx *ministore.Ministore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSpuUpSpuDelisting, bytes.NewReader(payload), "application/json;charset=utf-8")
}

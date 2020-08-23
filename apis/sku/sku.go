// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package sku SKU接口(修改需要重新上架商品)
package sku

import (
	"bytes"
	"github.com/fastwego/ministore"
)

const (
	apiAddSku      = "/product/sku/add"
	apiBatchAddSku = "/product/sku/batch_add"
	apiDelSku      = "/product/sku/del"
	apiGetSku      = "/product/sku/get"
	apiUpSku       = "/product/sku/update"
	apiUpSkuPrice  = "/product/sku/update_price"
	apiUpStock     = "/product/sku/stock/update"
	apiGetStock    = "/product/sku/stock/get"
)

/*
添加SKU

See: https://developers.weixin.qq.com/doc/MiniStore/minishopopencomponent/API/sku/add_sku.html
POST https://api.weixin.qq.com/product/sku/add?access_token=xxxxxxxxx
*/
func AddSku(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddSku, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量添加SKU

See: https://developers.weixin.qq.com/doc/MiniStore/minishopopencomponent/API/sku/batch_add_sku.html
POST https://api.weixin.qq.com/product/sku/batch_add?access_token=xxxxxxxxx
*/
func BatchAddSku(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBatchAddSku, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除SKU

See: https://developers.weixin.qq.com/doc/MiniStore/minishopopencomponent/API/sku/del_sku.html
POST https://api.weixin.qq.com/product/sku/del?access_token=xxxxxxxxx
*/
func DelSku(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelSku, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取SKU信息

See: https://developers.weixin.qq.com/doc/MiniStore/minishopopencomponent/API/sku/get_sku.html
POST https://api.weixin.qq.com/product/sku/get?access_token=xxxxxxxxx
*/
func GetSku(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetSku, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新SKU

See: https://developers.weixin.qq.com/doc/MiniStore/minishopopencomponent/API/sku/up_sku.html
POST https://api.weixin.qq.com/product/sku/update?access_token=xxxxxxxxx
*/
func UpSku(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpSku, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新SKU价格

See: https://developers.weixin.qq.com/doc/MiniStore/minishopopencomponent/API/sku/up_sku_price.html
POST https://api.weixin.qq.com/product/sku/update_price?access_token=xxxxxxxxx
*/
func UpSkuPrice(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpSkuPrice, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新库存

See: https://developers.weixin.qq.com/doc/MiniStore/minishopopencomponent/API/sku/up_stock.html
POST https://api.weixin.qq.com/product/sku/stock/update?access_token=xxxxxxxxx
*/
func UpStock(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpStock, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取库存

See: https://developers.weixin.qq.com/doc/MiniStore/minishopopencomponent/API/sku/get_stock.html
POST https://api.weixin.qq.com/product/sku/stock/get?access_token=xxxxxxxxx
*/
func GetStock(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetStock, bytes.NewReader(payload), "application/json;charset=utf-8")
}

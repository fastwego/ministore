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

// Package spu SPU接口(修改需要重新上架商品)
package spu

import (
	"bytes"
	"github.com/fastwego/ministore"
)

const (
	apiAddSpu         = "/product/spu/add"
	apiDelSpu         = "/product/spu/del"
	apiGetSpu         = "/product/spu/get"
	apiGetSpuList     = "/product/spu/get_list"
	apiUpSpu          = "/product/spu/update"
	apiUpSpuListing   = "/product/spu/listing"
	apiUpSpuDelisting = "/product/spu/delisting"
)

/*
添加商品

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/add_spu.html
POST https://api.weixin.qq.com/product/spu/add?access_token=xxxxxxxxx
*/
func AddSpu(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddSpu, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除商品

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/del_spu.html
POST https://api.weixin.qq.com/product/spu/del?access_token=xxxxxxxxx
*/
func DelSpu(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelSpu, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取商品

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/get_spu.html
POST https://api.weixin.qq.com/product/spu/get?access_token=xxxxxxxxx
*/
func GetSpu(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetSpu, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取商品列表

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/get_spu_list.html
POST https://api.weixin.qq.com/product/spu/get_list?access_token=xxxxxxxxx
*/
func GetSpuList(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetSpuList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新商品

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/up_spu.html
POST https://api.weixin.qq.com/product/spu/update?access_token=xxxxxxxxx
*/
func UpSpu(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpSpu, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
上架商品

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/up_spu_listing.html
POST https://api.weixin.qq.com/product/spu/listing?access_token=xxxxxxxxx
*/
func UpSpuListing(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpSpuListing, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
下架商品

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/up_spu_delisting.html
POST https://api.weixin.qq.com/product/spu/delisting?access_token=xxxxxxxxx
*/
func UpSpuDelisting(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpSpuDelisting, bytes.NewReader(payload), "application/json;charset=utf-8")
}

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

// Package cat 接入商品前必需接口
package cat

import (
	"bytes"
	"github.com/fastwego/ministore"
)

const (
	apiGetCatList         = "/product/category/get"
	apiGetBrand           = "/product/brand/get"
	apiGetFreightTemplate = "/product/delivery/get_freight_template"
)

/*
获取类目详情

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/cat/get_cat_list.html
POST https://api.weixin.qq.com/product/category/get?access_token=xxxxxxxxx
*/
func GetCatList(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetCatList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取品牌列表

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/cat/get_brand.html
POST https://api.weixin.qq.com/product/brand/get?access_token=xxxxxxxxx
*/
func GetBrand(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetBrand, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取运费模板

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/cat/get_freight_template.html
POST https://api.weixin.qq.com/product/delivery/get_freight_template?access_token=xxxxxxxxx
*/
func GetFreightTemplate(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetFreightTemplate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

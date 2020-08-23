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

// Package service 服务商接口
package service

import (
	"bytes"
	"github.com/fastwego/ministore"
)

const (
	apiCheckAuth           = "/product/service/check_auth"
	apiGetServiceList      = "/product/service/get_list"
	apiGetServiceOrderList = "/product/service/get_order_list"
)

/*
登录验证

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/service/check_auth.html
POST https://api.weixin.qq.com/product/service/check_auth?component_access_token=xxxxxxxxx
*/
func CheckAuth(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCheckAuth, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户购买的在有效期内的服务列表

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/service/get_service_list.html
POST https://api.weixin.qq.com/product/service/get_list?access_token=xxxxxxxxx
*/
func GetServiceList(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetServiceList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户购买的服务列表

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/service/get_service_order_list.html
POST https://api.weixin.qq.com/product/service/get_order_list?access_token=xxxxxxxxx
*/
func GetServiceOrderList(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetServiceOrderList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

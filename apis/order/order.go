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

// Package order 订单接口
package order

import (
	"bytes"
	"github.com/fastwego/ministore"
)

const (
	apiGetOrderList   = "/product/order/get_list"
	apiGetOrderDetail = "/product/order/get"
)

/*
获取订单列表

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/order/get_order_list.html
POST https://api.weixin.qq.com/product/order/get_list?access_token=xxxxxxxxx
*/
func GetOrderList(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetOrderList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取订单详情

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/order/get_order_detail.html
POST https://api.weixin.qq.com/product/order/get?access_token=xxxxxxxxx
*/
func GetOrderDetail(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetOrderDetail, bytes.NewReader(payload), "application/json;charset=utf-8")
}

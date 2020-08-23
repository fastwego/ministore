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

// Package coupon 优惠券接口
package coupon

import (
	"bytes"
	"github.com/fastwego/ministore"
)

const (
	apiGetCoupon  = "/product/coupon/get_list"
	apiPushCoupon = "/product/coupon/push"
)

/*
获取优惠券列表

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/coupon/get_coupon.html
POST https://api.weixin.qq.com/product/coupon/get_list?access_token=xxxxxxxxx
*/
func GetCoupon(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetCoupon, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
发放优惠券

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/coupon/push_coupon.html
POST https://api.weixin.qq.com/product/coupon/push?access_token=xxxxxxxxx
*/
func PushCoupon(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiPushCoupon, bytes.NewReader(payload), "application/json;charset=utf-8")
}

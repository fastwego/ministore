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

// Package delivery 物流接口
package delivery

import (
	"bytes"
	"github.com/fastwego/ministore"
)

const (
	apiGetDeliveryCompanyList = "/product/delivery/get_company_list"
	apiSendDelivery           = "/product/delivery/send"
)

/*
获取快递公司列表

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/delivery/get_delivery_company_list.html
POST https://api.weixin.qq.com/product/delivery/get_company_list?access_token=xxxxxxxxx
*/
func GetDeliveryCompanyList(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetDeliveryCompanyList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
订单发货

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/delivery/send_delivery.html
POST https://api.weixin.qq.com/product/delivery/send?access_token=xxxxxxxxx
*/
func SendDelivery(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSendDelivery, bytes.NewReader(payload), "application/json;charset=utf-8")
}

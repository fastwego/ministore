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

// Package store 店铺接口
package store

import (
	"bytes"
	"github.com/fastwego/ministore"
)

const (
	apiGetStoreInfo = "/product/store/get_info"
)

/*
获取基本信息

See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/store/get_store_info.html
POST https://api.weixin.qq.com/product/store/get_info?access_token=xxxxxxxxx
*/
func GetStoreInfo(ctx *ministore.MiniStore, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetStoreInfo, bytes.NewReader(payload), "application/json;charset=utf-8")
}

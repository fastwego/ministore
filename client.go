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

package ministore

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	Ctx *MiniStore
}

/*
微信 api 服务器地址
*/
var WXServerUrl = "https://api.weixin.qq.com"

//HTTPPost POST 请求
func (client *Client) HTTPPost(uri string, payload io.Reader, contentType string) (resp []byte, err error) {
	uri, err = client.applyAccessToken(uri)
	if err != nil {
		return
	}
	if client.Ctx.Logger != nil {
		client.Ctx.Logger.Printf("POST %s", uri)
	}
	response, err := http.Post(WXServerUrl+uri, contentType, payload)
	if err != nil {
		return
	}
	defer response.Body.Close()
	return responseFilter(response)
}

/*
在请求地址上附加上 access_token
*/
func (client *Client) applyAccessToken(oldUrl string) (newUrl string, err error) {
	accessToken := client.Ctx.Config.AccessToken
	fmt.Println(accessToken)

	if strings.Contains(oldUrl, "?") {
		newUrl = oldUrl + "&access_token=" + accessToken
	} else {
		newUrl = oldUrl + "?access_token=" + accessToken
	}
	return
}

/*
筛查微信 api 服务器响应，判断以下错误：
- http 状态码 不为 200
- 接口响应错误码 errcode 不为 0
*/
func responseFilter(response *http.Response) (resp []byte, err error) {
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("Status %s", response.Status)
		return
	}

	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	errorResponse := struct {
		Errcode int64  `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}{}
	err = json.Unmarshal(resp, &errorResponse)
	if err != nil {
		return
	}

	if errorResponse.Errcode != 0 {
		err = errors.New(string(resp))
		return
	}

	return
}

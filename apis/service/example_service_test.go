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

package service_test

import (
	"fmt"
	"github.com/fastwego/ministore"
	"github.com/fastwego/ministore/apis/service"
)

func ExampleServiceCheckAuth() {
	var ctx *ministore.MiniStore
	payload := []byte("{}")
	resp, err := service.CheckAuth(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleServiceGetServiceList() {
	var ctx *ministore.MiniStore
	payload := []byte("{}")
	resp, err := service.GetServiceList(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleServiceGetServiceOrderList() {
	var ctx *ministore.MiniStore
	payload := []byte("{}")
	resp, err := service.GetServiceOrderList(ctx, payload)
	fmt.Println(resp, err)
}

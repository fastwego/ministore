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

package sku_test

import (
	"fmt"
	"github.com/fastwego/ministore"
	"github.com/fastwego/ministore/apis/sku"
)

func ExampleSkuAddSku() {
	var ctx *ministore.MiniStore
	payload := []byte("{}")
	resp, err := sku.AddSku(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSkuBatchAddSku() {
	var ctx *ministore.MiniStore
	payload := []byte("{}")
	resp, err := sku.BatchAddSku(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSkuDelSku() {
	var ctx *ministore.MiniStore
	payload := []byte("{}")
	resp, err := sku.DelSku(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSkuGetSku() {
	var ctx *ministore.MiniStore
	payload := []byte("{}")
	resp, err := sku.GetSku(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSkuUpSku() {
	var ctx *ministore.MiniStore
	payload := []byte("{}")
	resp, err := sku.UpSku(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSkuUpSkuPrice() {
	var ctx *ministore.MiniStore
	payload := []byte("{}")
	resp, err := sku.UpSkuPrice(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSkuUpStock() {
	var ctx *ministore.MiniStore
	payload := []byte("{}")
	resp, err := sku.UpStock(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSkuGetStock() {
	var ctx *ministore.MiniStore
	payload := []byte("{}")
	resp, err := sku.GetStock(ctx, payload)
	fmt.Println(resp, err)
}

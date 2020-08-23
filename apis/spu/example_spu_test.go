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

package spu_test

import (
	"fmt"
	"github.com/fastwego/ministore"
	"github.com/fastwego/ministore/apis/spu"
)

func ExampleSpuAddSpu() {
	var ctx *ministore.MiniStore
	payload := []byte("{}")
	resp, err := spu.AddSpu(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSpuDelSpu() {
	var ctx *ministore.MiniStore
	payload := []byte("{}")
	resp, err := spu.DelSpu(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSpuGetSpu() {
	var ctx *ministore.MiniStore
	payload := []byte("{}")
	resp, err := spu.GetSpu(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSpuGetSpuList() {
	var ctx *ministore.MiniStore
	payload := []byte("{}")
	resp, err := spu.GetSpuList(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSpuUpSpu() {
	var ctx *ministore.MiniStore
	payload := []byte("{}")
	resp, err := spu.UpSpu(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSpuUpSpuListing() {
	var ctx *ministore.MiniStore
	payload := []byte("{}")
	resp, err := spu.UpSpuListing(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSpuUpSpuDelisting() {
	var ctx *ministore.MiniStore
	payload := []byte("{}")
	resp, err := spu.UpSpuDelisting(ctx, payload)
	fmt.Println(resp, err)
}

// Copyright 2020 FastWeGo
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

package cat

import (
	"fmt"
	"github.com/fastwego/ministore"
	"github.com/fastwego/ministore/test"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	test.Setup()
	os.Exit(m.Run())
}

func TestGetCatList(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(apiGetCatList, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})
	type args struct {
		ctx     *ministore.MiniStore
		payload []byte
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockMiniStore}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := GetCatList(tt.args.ctx, tt.args.payload)
			fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("CatGetCatList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CatGetCatList() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestCatGetBrand(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(apiGetBrand, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})
	type args struct {
		ctx     *ministore.MiniStore
		payload []byte
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockMiniStore}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := GetBrand(tt.args.ctx, tt.args.payload)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("CatGetBrand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CatGetBrand() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestCatGetFreightTemplate(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(apiGetFreightTemplate, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})
	type args struct {
		ctx     *ministore.MiniStore
		payload []byte
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockMiniStore}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := GetFreightTemplate(tt.args.ctx, tt.args.payload)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("CatGetFreightTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CatGetFreightTemplate() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

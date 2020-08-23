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

package ministore_test

import (
	"fmt"
	"github.com/fastwego/ministore"
	"github.com/fastwego/ministore/apis/cat"
	"testing"
)

func TestNew(t *testing.T) {
	app := ministore.New(ministore.Config{
		AccessToken: "aaabbb123",
	})

	payload := []byte(`{}`)

	resp, err := cat.GetCatList(app, payload)
	fmt.Println(string(resp), err)
}

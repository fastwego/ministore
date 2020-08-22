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

package main

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const ServerUrl = `https://developers.weixin.qq.com`

var apiUniqMap = map[string]bool{}

func main() {

	file, err := ioutil.ReadFile("./data/ministore_doc_links.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	pattern := `href="(/doc/ministore/minishopopencomponent/API/.+/.+\.html)"`
	reg := regexp.MustCompile(pattern)
	matched := reg.FindAllStringSubmatch(string(file), -1)

	for _, match := range matched {

		link := ServerUrl + match[1]
		resp, err := http.Get(link)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		name := doc.Find("#docContent div h2").First().Text()
		name = strings.Trim(name, "# ")
		//fmt.Println(name)
		_NAME_ := name
		_DESCRIPTION_ := ""


		api := doc.Find("#docContent div div").First().Text()
		api = strings.Trim(api, "# ")
		api = strings.ReplaceAll(api,"\n", " ")
		//fmt.Println(api)

		m4 := regexp.MustCompile(`http请求方式.+(POST|GET|PUT).+(https://api\.weixin\.qq\.com\/product\/(service|category|brand|delivery|spu|sku|order|delivery|coupon|store)\/\S+)`).FindStringSubmatch(api)
		//fmt.Println(len(m4))
		if len(m4) != 4 {
			continue
		}

		apiUrl := m4[2]
		method := m4[1]
		_REQUEST_ := method + " " +apiUrl
		_SEE_ := link
		_FUNCNAME_ := ""

		r := regexp.MustCompile(`/doc/ministore/minishopopencomponent/API/(.+)/(.+)\.html`)
		found := r.FindStringSubmatch(link)
		if len(found) == 3 {
			f1 := strcase.ToCamel(found[1])
			f2 := strings.ReplaceAll(found[2], "_", "")
			_FUNCNAME_ = f1 + strcase.ToCamel(f2)
		}

		tpl := strings.ReplaceAll(itemTpl, "_NAME_", _NAME_)
		tpl = strings.ReplaceAll(tpl, "_DESCRIPTION_", _DESCRIPTION_)
		tpl = strings.ReplaceAll(tpl, "_REQUEST_", _REQUEST_)
		tpl = strings.ReplaceAll(tpl, "_SEE_", _SEE_)
		tpl = strings.ReplaceAll(tpl, "_FUNCNAME_", _FUNCNAME_)

		fmt.Println(tpl)
	}

}

var itemTpl = `{
	Name:        "_NAME_",
	Description: "_DESCRIPTION_",
	Request:     "_REQUEST_",
	See:         "_SEE_",
	FuncName:    "_FUNCNAME_",
},`
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
	"log"
	"os"
)

/*
MiniStore 实例
*/
type MiniStore struct {
	Config Config
	Client Client
	Logger *log.Logger
}

/*
小商店配置
*/
type Config struct {
	AccessToken string
}

/*
创建小商店实例
*/
func New(config Config) (ministore *MiniStore) {
	instance := MiniStore{
		Config: config,
	}

	instance.Client = Client{Ctx: &instance}

	instance.Logger = log.New(os.Stdout, "[ministore] ", log.LstdFlags|log.Llongfile)

	return &instance
}

/*
SetLogger 日志记录 默认输出到 os.Stdout
可以新建 logger 输出到指定文件
如果不想开启日志，可以输出到 /dev/null log.SetOutput(ioutil.Discard)
*/
func (ministore *MiniStore) SetLogger(logger *log.Logger) {
	ministore.Logger = logger
}

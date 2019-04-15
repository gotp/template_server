/*
 * Copyright 2019 gotp
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"testing"
)

func TestInit_Nornmal_Case01(gtest *testing.T) {
	confFilePath := "./config_manager_test.conf"
	confMgr := GetInstance()

	ret := confMgr.Init(confFilePath)
	if ret != true {
		gtest.Fatal("Init failed!")
	}
	if confMgr.ConfigFilePath != confFilePath {
		gtest.Fatal("Init failed! ", confMgr.ConfigFilePath, " != ", confFilePath)
	}
	if confMgr.Addr != "127.0.0.1:8080" {
		gtest.Fatal("Init failed! ", confMgr.Addr, " != 127.0.0.1:8080")
	}
}

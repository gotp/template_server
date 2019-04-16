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

package module

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"regexp"

	glog "github.com/golang/glog"
)

// const
const serverConfigPath = "../conf/template_server.conf"

// internal
var configManager ConfigManager

func GetConfigManager() *ConfigManager {
	return &configManager
}

// ConfigManager struct & interface define
type ConfigManagerInterface interface {
	Init(configFilePath string) bool
}

type ConfigManager struct {
	ConfigFilePath      string
	Addr                string // ip:port
	PemPath             string // https pem file
	KeyPath             string // https key file
	RouterTableFilePath string
}

func (this *ConfigManager) Init(configFilePath string) bool {
	this.ConfigFilePath = configFilePath

	content, e := ioutil.ReadFile(this.ConfigFilePath)
	if e != nil {
		glog.Errorf("File error: %v\n", e)
		return false
	}

	content, e = this.stripJsonComments(content)
	if e != nil {
		glog.Errorf("File error: %v\n", e)
		return false
	}

	glog.Infof("Read config file: %s\n", this.ConfigFilePath)
	json.Unmarshal(content, &this)
	glog.Infof("Results: %v\n", this)

	return true
}

func (this *ConfigManager) stripJsonComments(data []byte) ([]byte, error) {
	data = bytes.Replace(data, []byte("\r"), []byte(""), 0) // Windows
	lines := bytes.Split(data, []byte("\n"))                //split to muli lines
	filtered := make([][]byte, 0)

	for _, line := range lines {
		match, err := regexp.Match(`^\s*#`, line)
		if err != nil {
			return nil, err
		}
		if !match {
			filtered = append(filtered, line)
		}
	}

	return bytes.Join(filtered, []byte("\n")), nil
}

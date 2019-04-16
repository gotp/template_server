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

package main

import (
	"flag"
	glog "github.com/golang/glog"
	framework "github.com/gotp/template_server/framework"
	module "github.com/gotp/template_server/module"
	service "github.com/gotp/template_server/service"
)

const (
	defaultConfigFilePath string = "./conf/template_server.conf"
)

var (
	configManager *module.ConfigManager
	server *framework.FrameworkServer
)

func init() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", defaultConfigFilePath, "Config file path")
	flag.Parse()

	// Load config
	configManager = module.GetConfigManager()
	if !configManager.Init(configFilePath) {
		glog.Fatal("Load server config failed!")
	}
	glog.Info("Load server config success")

	// Init framework 
	if !framework.Init() {
		glog.Fatal("Init framework failed!")
	}
	glog.Info("Init framework success")

	/*
	   	// Load router table
	   	routerTable := config.GetRouterTable()
	   	if !routerTable.Init(configManager.RouterTableFilePath) {
	   		glog.Fatal("Load router table failed!")
	   	}
	       glog.Info("Load router table success")

	        // Load resolver config
	       if svrresolver.GetResolverConfig().Init(configManager.ResolverConfigFilePath) == false {
	           glog.Fatal("Load resolver config failed!")
	       }
	       glog.Info("Load resolver config success")
	*/
}

func main() {
	// Get server
	server = framework.GetFrameworkServer()
	// Register service
	glog.Info("Register service...")
	service.RegisterService()
	// Start server
	glog.Info("Start server...")
	server.Start()
}

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
	proto "github.com/gotp/proto/template_server"
	rpc "github.com/gotp/template_server/framework/rpc"
	confmgr "github.com/gotp/template_server/module/config"
	service "github.com/gotp/template_server/service"

	"net/http"
	"strings"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const (
	defaultConfigFilePath string = "./conf/template_server.conf"
)

var configManager *confmgr.ConfigManager

func init() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", defaultConfigFilePath, "Config file path")
	flag.Parse()

	// Load config
	configManager = confmgr.GetInstance()
	if !configManager.Init(configFilePath) {
		glog.Fatal("Load server config failed!")
	}
	glog.Info("Load server config success")

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

func registerService() {
	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!! WARNING !!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	// Please not remove or modify comment below, it's anchor for new code
	// ############################# SERVICE #############################
	//proto.RegisterTemplateServiceServer(rpcServer, &service.TemplateService{})
}

func main() {
	// Create servers
	//httpHandler := httpservice.CreateHandler()
	rpcServer := rpc.NewServer()
	// Register service
	proto.RegisterTemplateServiceServer(rpcServer, &service.TemplateService{})
	// Start server
	glog.Info("Start server...")
	http.ListenAndServe(configManager.Addr,
		h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				glog.V(2).Info("Get a rpc request")
				rpcServer.ServeHTTP(w, r)
			} else {
				glog.V(2).Info("Get a http request")
				//httpHandler.ServeHTTP(w, r)
			}
		}), &http2.Server{}),
	)
}

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

package framework

import (
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
    "google.golang.org/grpc"
	glog "github.com/golang/glog"
	"net/http"
	"strings"
	"crypto/tls"
)

var (
	server FrameworkServer
)

// Framework server
type FrameworkServer struct {
	config *FrameworkConfig
	// rpc
	rpcServer  *grpc.Server
	// http
	httpServer  *http.Server
	httpHandler *http.ServeMux
}

func GetFrameworkServer() (*FrameworkServer) {
	return &server
}

func Init() bool {
	// load config
	server.config = NewFrameworkConfig()
	if !server.config.Init() {
		glog.Fatal("Load framework config failed!")
		return false
	}
	// new rpc server
	server.rpcServer = grpc.NewServer()
	// new http server
	server.httpHandler = http.NewServeMux()
	server.httpServer = &http.Server{
		Addr:    server.config.Addr,
		Handler: server.httpHandler,
		// Disable HTTP/2.
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
	return true
}

func (this *FrameworkServer) GetRpcServer() (*grpc.Server) {
	return this.rpcServer
}

func (this *FrameworkServer) RegisterHttpHandler(name string, function http.HandlerFunc) {
	this.httpHandler.HandleFunc(name, function)
}

func (this *FrameworkServer) Start() {
	http.ListenAndServe(this.config.Addr,
		h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				glog.V(2).Info("Get a rpc request")
				this.rpcServer.ServeHTTP(w, r)
			} else {
				glog.V(2).Info("Get a http request")
				this.httpHandler.ServeHTTP(w, r)
			}
		}), &http2.Server{}),
	)
}

func (this *FrameworkServer) StartHttpServer() {
	glog.Fatal(this.httpServer.ListenAndServe())
}

func (this *FrameworkServer) StartHttpsServer() {
	glog.Fatal(this.httpServer.ListenAndServeTLS(this.config.PemPath, this.config.KeyPath))
}

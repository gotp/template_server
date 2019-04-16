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
	"crypto/tls"
	"net/http"
	glog "github.com/golang/glog"
)

// Service handler register
var handlerFunc map[string]http.HandlerFunc

func RegisterServiceHandler(name string, function http.HandlerFunc) {
	if handlerFunc == nil {
		handlerFunc = make(map[string]http.HandlerFunc)
	}
	handlerFunc[name] = function
}

// Http server
type HttpServer struct {
	server  *http.Server
	handler *http.ServeMux
	handlerFunc map[string]http.HandlerFunc

	pemFile string
	keyFile string
}

func NewHttpServer(addr string, pem string, key string) *HttpServer {
	server := new(HttpServer)
	// build http handler
	server.handler = http.NewServeMux()
	for name, function := range handlerFunc {
		server.handler.HandleFunc(name, function)
	}
	// new http server
	server.server = &http.Server{
		Addr:    addr,
		Handler: server.handler,
		// Disable HTTP/2.
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
	// save pem & key
	server.pemFile = pem
	server.keyFile = key

	return server
}

func (this *HttpServer) RegisterServiceHandler(name string, function http.HandlerFunc) {
	if this.handlerFunc == nil {
		this.handlerFunc = make(map[string]http.HandlerFunc)
	}
	this.handlerFunc[name] = function
}

func (this *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	this.handler.ServeHTTP(w, r)
}

func (this *HttpServer) Start() {
	if this.pemFile != "" && this.keyFile != "" {
		this.StartHttpsServer()
	} else {
		this.StartHttpServer()
	}
}

func (this *HttpServer) StartHttpServer() {
	glog.Fatal(this.server.ListenAndServe())
}

func (this *HttpServer) StartHttpsServer() {
	glog.Fatal(this.server.ListenAndServeTLS(this.pemFile, this.keyFile))
}
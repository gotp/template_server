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

package http

import (
	"crypto/tls"
	"log"
	"net/http"

	config "github.com/gotp/template_server/module/config"
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
}

func NewServer() *HttpServer {
	server := new(HttpServer)
	// Init service
	InitRelayService()
	// build http handler
	server.handler = http.NewServeMux()
	for name, function := range handlerFunc {
		server.handler.HandleFunc(name, function)
	}
	// new http server
	server.server = &http.Server{
		Addr:    config.GetConfigManager().Addr,
		Handler: server.handler,
		// Disable HTTP/2.
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
	return server
}

func (this *HttpServer) Start() {
	configManager := config.GetConfigManager()
	if configManager.PemPath != "" && configManager.KeyPath != "" {
		this.StartHttpsServer(configManager.PemPath, configManager.KeyPath)
	} else {
		this.StartHttpServer()
	}
}

func (this *HttpServer) StartHttpServer() {
	log.Fatal(this.server.ListenAndServe())
}

func (this *HttpServer) StartHttpsServer(pemPath string, keyPath string) {
	log.Fatal(this.server.ListenAndServeTLS(pemPath, keyPath))
}

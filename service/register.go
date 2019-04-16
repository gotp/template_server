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

package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/golang/protobuf/jsonpb"
	proto "github.com/gotp/proto/template_server"
	framework "github.com/gotp/template_server/framework"
)

var (
	jsonMarshaler = jsonpb.Marshaler{EmitDefaults: true}
)

func RegisterService() {
	server := framework.GetFrameworkServer()
	rpcServer := server.GetRpcServer()
	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!! WARNING !!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	// Please not remove or modify comment below, it's anchor for new code
	// ############################# SERVICE #############################
	proto.RegisterTemplateServiceServer(rpcServer, &TemplateService{})

	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!! WARNING !!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	// Please not remove or modify comment below, it's anchor for new code
	// ############################# INTERFACE ###########################
	server.RegisterHttpHandler("",
		func(response http.ResponseWriter, request *http.Request) {
			var rpcRequest proto.TestRequest
			httpBody, _ := ioutil.ReadAll(request.Body)
			err := jsonpb.UnmarshalString(string(httpBody), &rpcRequest)
			if err != nil {
				fmt.Printf("%v", err)
				return
			}
			rpcResponse, err := new(TemplateService).Test(nil, &rpcRequest)
			if err != nil {
				fmt.Printf("%v", err)
				return
			}
			httpReponseBody, err := jsonMarshaler.MarshalToString(rpcResponse)
			if err != nil {
				fmt.Printf("%v", err)
				return
			}
			response.Write([]byte(httpReponseBody))
	})
}
package http_service

import (
	"fmt"
    "net/http"
    "io/ioutil"

    proto "online_consultant/proto/template_server"
    rpcservice "online_consultant/server/template_server/service/rpc_service"

    "github.com/golang/protobuf/jsonpb"
)

var jsonMarshaler = jsonpb.Marshaler{EmitDefaults: true}

func RegisterTemplateServiceHandler(httpServer *http.ServeMux) {
	/* 
	 * !!!!!!!!!!!! WARNING !!!!!!!!!!!!
	 * Please not remove or modify comment 
	 * below, it's anchor for new code
	 */
	// ############ INTERFACE ############
	httpServer.HandleFunc(GetHandlerUrl("/TemplateService/Test"), 
		func(response http.ResponseWriter, request *http.Request) {
			var rpcRequest proto.TestRequest

			httpBody, _ := ioutil.ReadAll(request.Body)

			err := jsonpb.UnmarshalString(string(httpBody), &rpcRequest)
			if err != nil {
				fmt.Printf("%v", err)
				return
			}

			rpcResponse, err := new(rpcservice.TemplateService).Test(nil, &rpcRequest)
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
		},
	)
}
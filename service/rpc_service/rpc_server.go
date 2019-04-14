package rpc_service

import (
    proto "online_consultant/proto/template_server"

    "google.golang.org/grpc"
)

func CreateServer() (*grpc.Server) {
    rpcServer := grpc.NewServer()
	registerRpcService(rpcServer)
	return rpcServer
}

func registerRpcService(rpcServer *grpc.Server) {
    /* 
	 * !!!!!!!!!!!! WARNING !!!!!!!!!!!!
	 * Please not remove or modify comment 
	 * below, it's anchor for new code
	 */
    // ############ SERVICE ############
    proto.RegisterTemplateServiceServer(rpcServer, &TemplateService{})
}

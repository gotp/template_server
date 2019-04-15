package template_service

import (
    "context"

    common "github.com/gotp/proto"
    proto "github.com/gotp/proto/template_server"
)

type TemplateService struct{}

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!! WARNING !!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// Please not remove or modify comment below, it's anchor for new code
// ############################ INTERFACE ############################
func (service *TemplateService) Test(ctx context.Context, in *proto.TestRequest) (*proto.TestResponse, error) {
    return &proto.TestResponse{
        Header: &common.ResponseHeader{
            Retcode: 0, 
            Retmsg: "ok",
            RequestId: "R0001",
        },
        Data: &proto.TestResponseData{
            Dummy: 1,
        },
    }, nil
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package hello

import (
	"context"
	
	"study_goframe/api/hello/v1"
	"study_goframe/api/hello/v2"
)

type IHelloV1 interface {
	Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error)
}

type IHelloV2 interface {
	Hello(ctx context.Context, req *v2.HelloReq) (res *v2.HelloRes, err error)
}



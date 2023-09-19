package user

import (
	"context"

	"github.com/gogf/gf/v2/net/gsvc"

	v1 "study_goframe/api/user/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedHServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterHServer(s.Server, &Controller{})
}

func (*Controller) HH(ctx context.Context, req *v1.Hello) (res *v1.Hello, err error) {
	gsvc.SetRegistry()
	grpcx.Resolver.Register()
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

package hello

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"study_goframe/api/hello/v2"
)

func (c *ControllerV2) Hello(ctx context.Context, req *v2.HelloReq) (res *v2.HelloRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

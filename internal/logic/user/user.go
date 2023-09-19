package user

import (
	"context"

	"study_goframe/internal/model"
)

func init() {
}

func New() *sConent {
	return &sConent{}
}

type sConent struct {
}

func (c *sConent) NoHello(ctx context.Context, in *model.HelloInput) error {
	return nil
}

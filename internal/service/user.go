// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"study_goframe/internal/model"
)

type (
	IConent interface {
		NoHello(ctx context.Context, in *model.HelloInput) error
	}
)

var (
	localConent IConent
)

func Conent() IConent {
	if localConent == nil {
		panic("implement not found for interface IConent, forgot register?")
	}
	return localConent
}

func RegisterConent(i IConent) {
	localConent = i
}

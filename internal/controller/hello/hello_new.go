// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package hello

import (
	"study_goframe/api/hello"
)

type ControllerV1 struct{}

func NewV1() hello.IHelloV1 {
	return &ControllerV1{}
}

type ControllerV2 struct{}

func NewV2() hello.IHelloV2 {
	return &ControllerV2{}
}


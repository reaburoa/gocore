package ecode

import (
	"github.com/spf13/cast"
)

var errorMap = map[string]int{}

// New 注册Error异常
func New(code int, err error) error {
	errorMap[err.Error()] = code
	return err
}

// Transform 通过error获取对应的code码
func Transform(err error) int {
	if err == nil {
		return -1
	}

	v, ok := errorMap[err.Error()]
	if !ok {
		return -1
	}
	return cast.ToInt(v)
}

package utils

import (
	"runtime"
	"strings"
)

// GetCurrentFunctionName 获取当前函数名称
func GetCurrentFunctionName() string {
	pc, _, _, _ := runtime.Caller(2)
	fullName := runtime.FuncForPC(pc).Name()
	parts := strings.Split(fullName, ".")
	return parts[len(parts)-1]
}

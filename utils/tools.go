package utils

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"google.golang.org/grpc/peer"
	"net"
	"net/url"
	"reflect"
	"strings"
)

// 辅助函数用于序列化错误信息
func MarshalError(code int, msg string) string {
	errorResponse := map[string]any{
		"code": code,
		"msg":  msg,
	}
	jsonData, _ := json.Marshal(errorResponse)
	return string(jsonData)
}

// 获取结构体实例中绑定了json标签的key值
func GetJSONKeysFromInstance(v interface{}) []string {
	var keys []string

	// 获取传入对象的值和类型
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	// 如果传入的是指针，获取指针指向的元素值和类型
	if typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	// 确保是结构体
	if typ.Kind() != reflect.Struct {
		return keys
	}

	// 遍历结构体的字段
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		valueField := val.Field(i)

		// 获取 json 标签
		jsonTag := field.Tag.Get("json")

		// 处理标签中的 "omitempty" 等情况，只取第一个逗号前的部分
		if jsonTag != "" && jsonTag != "-" {
			jsonKey := strings.Split(jsonTag, ",")[0]

			// 如果字段有值（非零值）才取出key
			if !valueField.IsZero() {
				keys = append(keys, jsonKey)
			}
		}
	}

	return keys
}

// 获取值不为""的字段名和字段值
func GetNonEmptyFields(v interface{}) ([]string, []string) {
	var fieldNames []string
	var fieldValues []string

	val := reflect.ValueOf(v)

	// 确保是结构体
	if val.Kind() == reflect.Ptr {
		val = val.Elem() // 解引用
	}
	if val.Kind() != reflect.Struct {
		return fieldNames, fieldValues
	}

	// 遍历结构体的字段
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i).Interface()

		// 检查字段值是否非空字符串
		if strValue, ok := value.(string); ok && strValue != "" {
			fieldNames = append(fieldNames, field.Name)
			fieldValues = append(fieldValues, strValue)
		}
	}

	return fieldNames, fieldValues
}

// RemoveString 从切片中移除指定的字符串
func RemoveString(slice []string, str string) []string {
	var result []string
	for _, item := range slice {
		if item != str {
			result = append(result, item)
		}
	}
	return result
}

// JSON 转换为表单数据
func JSONToFormData(jsonData interface{}) (string, error) {

	ks, vs := GetNonEmptyFields(jsonData)

	// 创建 url.Values 对象
	formData := url.Values{}

	// 遍历 JSON 数据，将其转换为表单数据
	for i := 0; i < len(ks); i++ {
		formData.Set(ks[i], fmt.Sprintf("%v", vs[i]))
	}

	// 编码为表单格式
	return formData.Encode(), nil
}

func GetIPToCtx(c context.Context) (string, error) {
	// 获取客户端的 IP 地址
	p, ok := peer.FromContext(c)
	// 提取 IP 部分（去掉端口）
	host, _, err := net.SplitHostPort(p.Addr.String())
	if err != nil {
		return "", err
	}
	if !ok {
		return "", err
	}
	return host, nil
}

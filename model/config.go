package model

import (
	grpc_config "bigagent_server/grpcs/grpc_config"
	"encoding/json"
	"fmt"
)

type AuthDetails interface {
	ApplyAuth(args ...interface{}) error
}

//type FieldMapping struct {
//	grpc_config.FieldMapping
//	//StructField string `json:"struct_field"`
//	//Type        string `json:"type"`
//}
//
//type NetworkInfo struct {
//	grpc_config.NetworkInfo
//	//Protocol string `json:"protocol"`
//	//Host     string `json:"host"`
//	//Port     int    `json:"port"`
//	//Path     string `json:"path"`
//}

type AgentConfig struct {
	AuthName     string                               `json:"auth_name"`
	AuthDetails  AuthDetails                          `json:"auth_details"`
	FieldMapping map[string]*grpc_config.FieldMapping `json:"field_mapping"`
	NetworkInfo  *grpc_config.NetworkInfo             `json:"network_info"`

	//下面为前端传入的各种认证数据
	Token string `json:"token"`
}

func (a *AgentConfig) UnmarshalAuthDetails(body []byte) error {
	var data map[string]interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	if v, ok := data["auth_name"]; ok {
		switch v.(string) {
		case "tokenauth":
			t := new(TokenAuth)
			err := t.ApplyAuth(a.Token)
			if err != nil {
				return err
			}
			a.AuthDetails = t
		case "xxx":
		// ...
		default:
			return fmt.Errorf("无效的认证模式%s", v)
		}
	}
	return err
}

// 接收案例
//{
//"auth_mode": "token",
//"auth_details": {
//"token": "my-secret-token"
//},
//"field_mapping": {
//"cmdb_field1": { "struct_field": "Name", "type": "string" },
//"cmdb_field2": { "struct_field": "Age", "type": "integer" }
//},
//"data_rules": {
//"generate_data": true,
//"fields": {
//"Name": { "default": "Default Name" },
//"Age": { "default": 30 }
//}
//},
//"network_info": {
//"protocol": "http",
//"host": "cmdb.example.com",
//"port": 8080,
//"path": "/api/v1/data"
//}
//}

package mysqldb

import (
	"bigagent_server/model"
	"fmt"
	"reflect"
	"testing"
)

func TestAgentSelect(t *testing.T) {
	// 准备测试数据
	testUUID := "test-uuid-12345"
	//testData := model.AgentInfo{
	//	UUID:         testUUID,
	//	NetIP:        "192.168.1.100",
	//	Hostname:     "test-host",
	//	IPv4First:    "192.168.1.1",
	//	Active:       1,
	//	Status:       "running",
	//	ActionDetail: "none",
	//}
	// 检查数据库中是否已删除
	agentSelect, err := AgentSelect(testUUID)
	fmt.Println(agentSelect, err)
}

func TestAgentInfoSelectByKeys(t *testing.T) {
	type args struct {
		cp   string
		ps   string
		uuid string
		ip   string
		t    string
		p    string
		a    string
		c    string
	}
	tests := []struct {
		name    string
		args    args
		want    []model.AgentInfo
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				cp: "1",
				ps: "10",
				c:  "25",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AgentInfoSelectByKeys(tt.args.cp, tt.args.ps, tt.args.uuid, tt.args.ip, tt.args.t, tt.args.p, tt.args.a, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentInfoSelectByKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AgentInfoSelectByKeys() got = %v, want %v", got, tt.want)
			}
		})
	}
}

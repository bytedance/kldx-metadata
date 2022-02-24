package metadata

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/kldx-metadata/object/fields"
	"testing"
)

func TestGetFields(t *testing.T) {
	// 定义对应的结构，注意 tag 对应 apiName
	type Result struct {
		CreatedBy fields.Lookup `json:"_createdBy"`
	}
	// 调用接口，decode 拿到结果
	res := Result{}
	err := Object("_user").GetFields(&res)
	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}

func TestGetField(t *testing.T) {
	res := fields.Multilingual{}
	err := Object("_user").GetField("_createdBy", &res)
	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}

package goapi_metadata

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/kldx-metadata/object/fields"
	"testing"
)

func TestGetFields(t *testing.T) {
	res := map[string]interface{}{}
	err := Metadata.Object("allTypeFieldObject").GetFields(&res)
	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}

func TestGetField(t *testing.T) {
	res := fields.Lookup{}
	err := Metadata.Object("allTypeFieldObject").GetField("lookupType", &res)
	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}
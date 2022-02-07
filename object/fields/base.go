package fields

import "github.com/bytedance/kldx-metadata/structs"

type FieldBase struct {
	Type    string        `json:"type"`
	ApiName string        `json:"apiName"`
	Label   structs.I18ns `json:"label"`
}

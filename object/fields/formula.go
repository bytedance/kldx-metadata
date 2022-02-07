package fields

import "github.com/bytedance/kldx-metadata/structs"

type Formula struct {
	FieldBase
	ReturnType string        `json:"returnType"`
	Formula    structs.I18ns `json:"formula"`
}

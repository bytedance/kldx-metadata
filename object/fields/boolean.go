package fields

import "github.com/bytedance/kldx-metadata/structs"

type Boolean struct {
	FieldBase
	DescriptionWhenTrue  structs.I18ns `json:"descriptionWhenTrue"`
	DescriptionWhenFalse structs.I18ns `json:"descriptionWhenFalse"`
	DefaultValue         bool          `json:"defaultValue"`
}

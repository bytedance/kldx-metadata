package fields

import "github.com/bytedance/kldx-metadata/structs"

type Option struct {
	FieldBase
	Required            bool          `json:"required"`
	Multiple            bool          `json:"multiple"`
	OptionSource        string        `json:"optionSource"` // 暴露给开发者：custom, global; 底层存储：local,  global
	GlobalOptionApiName string        `json:"globalOptionApiName"`
	OptionList          []*OptionItem `json:"optionList"`
}

type OptionItem struct {
	Label       structs.I18ns `json:"label"`
	ApiName     string        `json:"apiName"`
	Description structs.I18ns `json:"description"`
	Color       string        `json:"color"`
	Active      bool          `json:"active"`
}

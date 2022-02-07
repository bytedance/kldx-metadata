package fields

import "github.com/bytedance/kldx-metadata/structs"

type ExtractSingleRecord struct {
	FieldBase
	CompositeTypeApiName string                 `json:"compositeTypeApiName"`
	SubFields            map[string]interface{} `json:"subFields"`
	Filter               *structs.Criterion     `json:"filter"`
	SortConditions       *structs.Sorts         `json:"sortConditions"`
	RecordPosition       int64                  `json:"recordPosition"`
}

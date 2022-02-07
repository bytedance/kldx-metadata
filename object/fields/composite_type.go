package fields

type CompositeType struct {
	FieldBase
	CompositeTypeApiName string                 `json:"compositeTypeApiName"`
	Required             bool                   `json:"required"`
	Multiple             bool                   `json:"multiple"`
	SubFields            map[string]interface{} `json:"subFields"`
}

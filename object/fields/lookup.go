package fields

type SortCondition struct {
	FieldApiName string `json:"fieldApiName"`
	Sort         string `json:"sort"`//"asc"|"desc"
}

type Lookup struct {
	FieldBase
	Required       bool            `json:"required"`
	Multiple       bool            `json:"multiple"`
	ObjectApiName  string          `json:"objectApiName"`
	Hierarchy      bool            `json:"hierarchy"`
	DisplayStyle   string          `json:"displayStyle"`
	SortConditions []SortCondition `json:"sortConditions"`
	Filter         []interface{}   `json:"filter"` // TODO 未返回
}

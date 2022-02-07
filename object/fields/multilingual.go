package fields

type Multilingual struct {
	FieldBase
	Required      bool  `json:"required"`
	Unique        bool  `json:"unique"`
	CaseSensitive bool  `json:"caseSensitive"`
	Multiline     bool  `json:"multiline"`
	MaxLength     int64 `json:"maxLength"`
}

package fields

type Text struct {
	FieldBase
	Required        bool   `json:"required"`
	Unique          bool   `json:"unique"`
	CaseSensitive   bool   `json:"caseSensitive"`
	Multiline       bool   `json:"multiline"`
	MaxLength       int64  `json:"maxLength"`
	ValidationRegex string `json:"validationRegex"`
	ErrorMsg        string `json:"errorMsg"`
}

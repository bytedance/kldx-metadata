package fields

type RichText struct {
	FieldBase
	Required  bool `json:"required"`  // 默认值：false
	MaxLength int  `json:"maxLength"` // 默认值：1000
}

package fields

type AutoID struct {
	FieldBase
	GenerateMethod string `json:"generateMethod"` // "random"|"incremental"
	DigitsNumber   int64  `json:"digitsNumber"`
	Prefix         string `json:"prefix"`
	Suffix         string `json:"suffix"`
}

package fields

type Email struct {
	FieldBase
	Required bool `json:"required"`
	Unique   bool `json:"unique"`
}

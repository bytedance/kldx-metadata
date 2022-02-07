package fields

type File struct {
	FieldBase
	Required  bool     `json:"required"`
	Multiple  bool     `json:"multiple"`
	FileTypes []string `json:"fileTypes"`
}

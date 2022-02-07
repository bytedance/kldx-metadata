package structs

type I18n struct {
	LanguageCode int64  `mapstructure:"language_code" json:"language_code" yaml:"LanguageCode"`
	Text         string `json:"text" yaml:"Text"`
}

type I18ns []I18n

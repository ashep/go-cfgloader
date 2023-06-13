package cfgloader_test

type outStruct struct {
	Foo string `yaml:"foo" json:"foo"`
	Bar string `yaml:"bar" json:"bar"`
}

var testSchema = []byte(`{
	"$schema": "https://json-schema.org/draft/2020-12/schema",
	"type": "object",
	"properties": {
		"foo": {
			"type": "string",
			"maxLength": 2
		},
		"bar": {
			"type": "string",
			"minLength": 1
		}
	}
}
`)

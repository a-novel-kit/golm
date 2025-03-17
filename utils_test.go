package golm_test

import "text/template"

var TestTemplate = template.Must(template.New("test").Parse("{{.Content}}"))

type TestTemplateData struct {
	Content string
}

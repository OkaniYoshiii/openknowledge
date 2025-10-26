package templates

import "html/template"

type Config struct {
	TemplateDir  string
	BaseTemplate *template.Template
}

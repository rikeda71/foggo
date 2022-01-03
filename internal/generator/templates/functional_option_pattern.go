package templates

const FOPTemplate = TemplateBase + `
type {{ .structName }}Option func(*{{ .structName }})

func New{{ .structName }}(options ...{{ .structName }}Option) *{{ .structName }} {
	s := &{{ .structName }}{}

	for _, option := range options {
		option(s)
	}

	return s
}
{{ range .fields }}
func With{{ .Name }}({{ .Name }} {{ .Type }}) {{ $.structName }}Option {
	return func(args *{{ $.structName }}) {
		args.{{ .Name }} = {{ .Name }}
	}
}
{{ end }}`

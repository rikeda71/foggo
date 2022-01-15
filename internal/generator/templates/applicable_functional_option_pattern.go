package templates

const AFOPTemplate = TemplateBase + `
type {{ .structName }}Option interface {
	apply(*{{ .structName }})
}
{{ range .fields }}{{ if ne .Ignore true}}
type {{ .Name }}Option struct {
	{{ .Name }} {{ .Type }}
}

func (o {{ .Name }}Option) apply(s {{ $.structName }}) {
	s.{{ .Name }} = o.{{ .Name }}
}
{{ end }}{{ end }}
func New{{ .structName }}(options ...{{ .structName }}Option) *{{ .structName }} {
	s := &{{ .structName }}{}

	for _, option := range options {
		option.apply(s)
	}

	return s
}
`

package generator

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"text/template"

	"github.com/s14t284/foggo/internal/generator/templates"
)

type Generator struct {
	goimports bool
}

func InitializeGenerator() *Generator {
	return &Generator{
		goimports: true,
	}
}

// GenerateFOP is the function to generate code of Functional Option Pattern from struct
func (g *Generator) GenerateFOP(pkgName string, structName string, sts []*StructField) (string, error) {
	tpl := template.Must(template.New("t").Parse(templates.FOPTemplate))
	return g.generateInternal(pkgName, structName, sts, tpl)
}

// GenerateAFOP is the function to generate code of Applicable Functional Option Pattern from struct
func (g *Generator) GenerateAFOP(pkgName string, structName string, sts []*StructField) (string, error) {
	tpl := template.Must(template.New("t").Parse(templates.AFOPTemplate))
	return g.generateInternal(pkgName, structName, sts, tpl)
}

// generateInternal is the function of the internal logic to generate Functional Option Pattern code
func (g *Generator) generateInternal(pkgName string, structName string, sts []*StructField, tpl *template.Template) (string, error) {
	if !g.checkStructFieldFormat(sts) {
		return "", fmt.Errorf("%s have same name fields", structName)
	}

	data := map[string]interface{}{
		"pkgName":    pkgName,
		"structName": structName,
		"fields":     sts,
	}
	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, data); err != nil {
		return "", fmt.Errorf("render code with template error: %w", err)
	}

	return g.format(buf)
}

func (g *Generator) format(b *bytes.Buffer) (string, error) {
	var err error
	code := b.String()
	if g.goimports {
		code, err = g.applyGoImports(code)
		if err != nil {
			return "", err
		}
	}
	return code, nil
}

func (g *Generator) applyGoImports(code string) (string, error) {
	cmd := exec.Command("goimports")
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		return "", fmt.Errorf("get stdin pipeline error in execution goimports: %w", err)
	}

	out, errOut := &bytes.Buffer{}, &bytes.Buffer{}
	cmd.Stdout = out
	cmd.Stderr = errOut

	err = cmd.Start()
	if err != nil {
		return "", fmt.Errorf("exec goimports error: %w", err)
	}

	_, err = io.WriteString(stdinPipe, code)
	if err != nil {
		return "", fmt.Errorf("write goimports result error: [error][%w][errOut][%s]", err, errOut.String())
	}
	err = stdinPipe.Close()
	if err != nil {
		return "", fmt.Errorf("close stdin pipeline error in execution goimports: %w", err)
	}

	err = cmd.Wait()
	if err != nil {
		return "", fmt.Errorf("write goimports result error: [error][%w][errOut][%s]", err, errOut.String())
	}

	return out.String(), nil
}

func (g *Generator) checkStructFieldFormat(sts []*StructField) bool {
	for _, st1 := range sts {
		name := strings.ToLower(st1.Name)
		cnt := 0
		for _, st2 := range sts {
			if name == strings.ToLower(st2.Name) {
				cnt++
			}
		}
		if cnt > 1 {
			return false
		}
	}
	return true
}

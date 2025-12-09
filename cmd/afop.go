/*
Copyright © 2025 rikeda71 rikeda71@gmail.com
*/
package cmd

import (
	"io"
	"os"
	"path"

	"github.com/rikeda71/foggo/internal/generator"
	"github.com/rikeda71/foggo/internal/logger"
	"github.com/rikeda71/foggo/internal/parser"
	"github.com/rikeda71/foggo/internal/writer"
	"github.com/spf13/cobra"
)

func initializeAfopCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "afop",
		Short: "command to generate 'Applicable Functional Option Pattern' code of golang",
		Long: `'afop' is the command to command to generate 'Applicable Functional Option Pattern' code of golang.
ref.
- https://github.com/uber-go/guide/blob/master/style.md#functional-options
- https://ww24.jp/2019/07/go-option-pattern(in Japanese)
`,
		RunE: func(_ *cobra.Command, _ []string) error {
			out := os.Stdout
			return generateAFOP(out)
		},
	}
}

// generateAFOP generate Applicable Functional Option Pattern code
func generateAFOP(out io.Writer) error {
	l := logger.InitializeLogger(out, "[AFOP Generator] ")
	g := generator.InitializeGenerator()
	w, err := writer.InitializeWriter(l)
	if err != nil {
		return err
	}

	p := Args.Package
	if p != "." {
		p = "./" + path.Clean(Args.Package)
	}
	pkg, err := parser.ParsePackageInfo(p)
	if err != nil {
		return err
	}

	fields, i, err := parser.CollectFields(Args.Struct, pkg.AstFiles)
	if err != nil {
		return err
	}

	code, err := g.GenerateAFOP(pkg.Name, Args.Struct, fields)
	if err != nil {
		return err
	}

	err = w.Write(code, pkg.Paths[i])
	if err != nil {
		return err
	}

	return nil
}

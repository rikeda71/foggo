/*
Copyright © 2022 s14t284 rikeda71@gmail.com

*/
package cmd

import (
	"io"
	"os"
	"path"

	"github.com/s14t284/foggo/internal/generator"
	"github.com/s14t284/foggo/internal/logger"
	"github.com/s14t284/foggo/internal/parser"
	"github.com/s14t284/foggo/internal/writer"
	"github.com/spf13/cobra"
)

func initializeAfocCommand() *cobra.Command {
	// afocCmd represents the afoc command
	return &cobra.Command{
		Use:   "afoc",
		Short: "command to generate 'Applicable Functional Option Pattern' code of golang",
		Long: `'afoc' is the command to command to generate 'Applicable Functional Option Pattern' code of golang.
ref.
- https://github.com/uber-go/guide/blob/master/style.md#functional-options
- https://ww24.jp/2019/07/go-option-pattern(in Japanese)
`,
		RunE: func(_ *cobra.Command, _ []string) error {
			out := os.Stdout
			return generateAFOC(out)
		},
	}
}

// generateAFOC generate functional option pattern code
func generateAFOC(out io.Writer) error {
	l := logger.InitializeLogger(out, "[AFOC Generator] ")
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

	code, err := g.GenerateFOP(pkg.Name, Args.Struct, fields)
	if err != nil {
		return err
	}

	err = w.Write(code, pkg.Paths[i])
	if err != nil {
		return err
	}

	return nil
}

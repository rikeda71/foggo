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

func initializeFocCommand() *cobra.Command {
	// focCmd represents the foc command
	return &cobra.Command{
		Use:   "foc",
		Short: "command to generate 'Functional Option Pattern' code of golang",
		Long: `'foc' is the command to command to generate 'Functional Option Pattern' code of golang.
ref.
- https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
- https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis`,
		RunE: func(_ *cobra.Command, _ []string) error {
			out := os.Stdout
			return generateFOC(out)
		},
	}
}

// generateFOC generate functional option pattern code
func generateFOC(out io.Writer) error {
	l := logger.InitializeLogger(out, "[FOC Generator] ")
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

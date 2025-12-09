// Package cmd provides CLI commands for foggo.
// Copyright © 2025 rikeda71 rikeda71@gmail.com
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

func initializeFopCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "fop",
		Short: "command to generate 'Functional Option Pattern' code of golang",
		Long: `'fop' is the command to command to generate 'Functional Option Pattern' code of golang.
ref.
- https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
- https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis`,
		RunE: func(_ *cobra.Command, _ []string) error {
			out := os.Stdout
			return generateFOP(out)
		},
	}
}

// generateFOP generate functional option pattern code
func generateFOP(out io.Writer) error {
	l := logger.InitializeLogger(out, "[FOP Generator] ")
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

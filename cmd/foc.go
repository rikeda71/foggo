/*
Copyright © 2022 s14t284 rikeda71@gmail.com

*/
package cmd

import (
	"github.com/s14t284/foggo/internal/generator"
	"github.com/s14t284/foggo/internal/parser"
	"github.com/s14t284/foggo/internal/writer"
	"github.com/spf13/cobra"
)

func generateFOC(_ *cobra.Command, _ []string) error {
	g := generator.InitializeGenerator()
	w := writer.InitializeWriter()

	pkg, err := parser.ParsePackageInfo(Flag.Package)
	if err != nil {
		return err
	}

	sts, i, err := parser.CollectFields(Flag.Source, pkg.AstFiles)
	if err != nil {
		return err
	}

	code, err := g.GenerateFOP(pkg.Name, Flag.Source, sts)
	if err != nil {
		return err
	}

	err = w.Write(code, pkg.Paths[i])
	if err != nil {
		return err
	}

	return nil
}

// focCmd represents the foc command
var focCmd = &cobra.Command{
	Use:   "foc",
	Short: "command to generate 'Functional Option Pattern' code of golang",
	Long: `'foc' is the command to command to generate 'Functional Option Pattern' code of golang.
ref.
- https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
- https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis`,
	RunE: generateFOC,
}

func init() {
	rootCmd.AddCommand(focCmd)
}

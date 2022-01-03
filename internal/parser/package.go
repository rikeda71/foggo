package parser

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/packages"
)

type PackageInfo struct {
	Name      string
	AstFiles  []*ast.File
	Paths     []string
	FileNames []string
}

func ParsePackageInfo(path string) (*PackageInfo, error) {
	cfg := &packages.Config{
		Mode:  packages.NeedName | packages.NeedSyntax | packages.NeedFiles,
		Tests: false,
	}
	packagesList, err := packages.Load(cfg, path)
	if err != nil {
		return nil, err
	}

	if len(packagesList) != 1 {
		return nil, fmt.Errorf("error: %d packages found", len(packagesList))
	}

	p := packagesList[0]
	return &PackageInfo{
		Name:     p.Name,
		AstFiles: p.Syntax,
		Paths:    p.GoFiles,
	}, nil
}

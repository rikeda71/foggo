package parser

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/fatih/structtag"
	"github.com/s14t284/foggo/internal/generator"
)

const (
	TagKey      = "foggo"
	IgnoreValue = "-"
)

// CollectFields is function to get fields of struct type and index of struct type from ast files
func CollectFields(source string, astFiles []*ast.File) ([]*generator.StructField, int, error) {
	for i, af := range astFiles {
		for _, decl := range af.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				structName := typeSpec.Name.Name
				if source != structName {
					continue
				}
				structType, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}

				converted, err := convertToStructFieldsFromAstFields(structType.Fields.List)
				if err != nil {
					return nil, -1, err
				}

				return converted, i, nil
			}
		}
	}
	return nil, -1, fmt.Errorf("there is no suitable struct that matches given source [source=%s]", source)
}

func convertToStructFieldsFromAstFields(fields []*ast.Field) ([]*generator.StructField, error) {
	sfs := make([]*generator.StructField, len(fields))
	for i, f := range fields {
		sf, err := internalConvert(f)
		if err != nil {
			return nil, fmt.Errorf("convert internal type error in %dth field: %w", i, err)
		}
		sfs[i] = sf
	}
	return sfs, nil
}

func internalConvert(field *ast.Field) (*generator.StructField, error) {
	tag, err := parseTag(field.Tag)
	if err != nil {
		return nil, err
	}
	ignore := tag != nil && tag.Name == IgnoreValue
	fieldType := types.ExprString(field.Type)

	var fieldName string
	if len(field.Names) > 0 {
		fieldName = field.Names[0].Name
	} else {
		// get Struct Name from 'package.Struct'
		chunks := strings.Split(fieldType, ".")
		fieldName = strings.TrimPrefix(chunks[len(chunks)-1], "*")
	}
	return &generator.StructField{
		Name:   fieldName,
		Type:   fieldType,
		Ignore: ignore,
	}, nil
}

func parseTag(tag *ast.BasicLit) (*structtag.Tag, error) {
	// not set tag
	if tag == nil {
		return nil, nil
	}

	tags, err := structtag.Parse(strings.Replace(tag.Value, "`", "", -1))
	if err != nil {
		return nil, fmt.Errorf("parse tag error: %w", err)
	}

	val, err := tags.Get(TagKey)
	if err != nil {
		return nil, fmt.Errorf("get tag error: %w", err)
	}

	return val, nil
}

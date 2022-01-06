package parser

import (
	"fmt"
	"go/ast"
	"os"
	"testing"

	"github.com/s14t284/foggo/internal/generator"
	"github.com/stretchr/testify/assert"
)

func TestCollectFields(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}
	st, err := ParsePackageInfo(cwd + "/../../testdata")
	if err != nil {
		t.FailNow()
	}

	type args struct {
		typeName string
		astFiles []*ast.File
	}
	tests := []struct {
		name    string
		args    args
		want    []*generator.StructField
		want1   int
		wantErr assert.ErrorAssertionFunc
	}{
		{"nominal_Data1", args{typeName: "Data1", astFiles: st.AstFiles}, []*generator.StructField{{Name: "A", Type: "string", Ignore: false}, {Name: "B", Type: "int", Ignore: false}, {Name: "C", Type: "string", Ignore: true}}, 0, assert.NoError},
		{"nominal_Data2", args{typeName: "Data2", astFiles: st.AstFiles}, []*generator.StructField{{Name: "A", Type: "string", Ignore: true}, {Name: "B", Type: "string", Ignore: false}, {Name: "C", Type: "string", Ignore: false}}, 1, assert.NoError},
		{"non_nominal: astFiles is empty", args{typeName: "", astFiles: []*ast.File{}}, nil, -1, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)
			got, got1, err := CollectFields(tt.args.typeName, tt.args.astFiles)
			if !tt.wantErr(t, err, fmt.Sprintf("CollectFields(%v, %v)", tt.args.typeName, tt.args.astFiles)) {
				return
			}
			a.Equal(tt.want, got)
			a.Equal(tt.want1, got1)
		})
	}
}

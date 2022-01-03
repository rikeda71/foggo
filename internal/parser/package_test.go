package parser_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/s14t284/foggo/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestParsePackageInfo1(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *parser.PackageInfo
		wantErr assert.ErrorAssertionFunc
	}{
		{"nominal", args{cwd + "/../../testdata"}, &parser.PackageInfo{Name: "testdata"}, assert.NoError},
		{"non_nominal: not found package", args{"./foo/bar"}, &parser.PackageInfo{Name: ""}, assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)
			got, err := parser.ParsePackageInfo(tt.args.path)
			if !tt.wantErr(t, err, fmt.Sprintf("ParsePackageInfo(%v)", tt.args.path)) {
				return
			}
			a.Equal(got.Name, tt.want.Name)
		})
	}
}

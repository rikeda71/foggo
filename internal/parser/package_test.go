package parser

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePackageInfo(t *testing.T) {
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
		want    *PackageInfo
		wantErr assert.ErrorAssertionFunc
	}{
		{"nominal", args{cwd + "/../../testdata"}, &PackageInfo{Name: "testdata"}, assert.NoError},
		{"non_nominal: not found package", args{"./foo/bar"}, &PackageInfo{Name: ""}, assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)
			got, err := ParsePackageInfo(tt.args.path)
			if !tt.wantErr(t, err, fmt.Sprintf("ParsePackageInfo(%v)", tt.args.path)) {
				return
			}
			a.Equal(got.Name, tt.want.Name)
		})
	}
}

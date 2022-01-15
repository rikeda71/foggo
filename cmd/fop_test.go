package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_initializeFopCommand(t *testing.T) {
	// assert initialization and check required parameters
	a := assert.New(t)
	cmd := initializeFopCommand()
	a.NotEqual("", cmd.Use)
	a.NotEqual("", cmd.Short)
	a.NotEqual("", cmd.Long)
	a.NotNil(cmd.RunE)
}

func Test_generateFOP(t *testing.T) {
	tests := []struct {
		name     string
		struct_  string
		package_ string
		wantOut  string
		wantErr  assert.ErrorAssertionFunc
	}{
		{"nominal: Data1", "Data1", "../testdata", "success to write functional option pattern code to", assert.NoError},
		{"nominal: Data2", "Data2", "../testdata", "success to write functional option pattern code to", assert.NoError},
		{"non_nominal: parse package error", "Data2", "./", "", assert.Error},
		{"non_nominal: collect fields from struct type error", "Data3", "../testdata", "", assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Args.Struct = tt.struct_
			Args.Package = tt.package_
			out := &bytes.Buffer{}
			err := generateFOP(out)
			if !tt.wantErr(t, err, fmt.Sprintf("generateFOP(%v)", out)) {
				return
			}
			assert.Containsf(t, out.String(), tt.wantOut, "generateFOP(%v)", out)

			// remove generated files
			files, err := filepath.Glob(fmt.Sprintf("%s/*_gen.go", tt.package_))
			assert.NoError(t, err)
			for _, f := range files {
				fmt.Println(f)
				err = os.Remove(f)
				assert.NoError(t, err)
			}
		})
	}
}

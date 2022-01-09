package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_initializeRootCmd(t *testing.T) {
	// assert initialization and check required parameters
	a := assert.New(t)
	cmd, err := initializeRootCmd()
	a.NotNil(err)
	a.NotEqual("", cmd.Use)
	a.NotEqual("", cmd.Short)
	a.NotEqual("", cmd.Long)
}

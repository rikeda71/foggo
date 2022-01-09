/*
Copyright © 2022 s14t284 rikeda71@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initializeRootCmd generate cobra.Command of root command
func initializeRootCmd() (*cobra.Command, error) {
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:   "foggo",
		Short: "'foggo' is the cli to generate 'Functional Option Pattern' code of golang from golang code",
		Long: `'foggo' is the cli to generate 'Functional Option Pattern' code of golang from golang code
# Example:

## Generate 'Functional Option Pattern' code
$ foggo foc --source ${STRUCT_NAME} --package /path/to/package
`,
	}

	// set arguments
	args := []string{"source", "package"}
	shortArgs := []string{"s", "p"}
	usages := []string{
		"Target struct name (required)",
		"Package name having target struct (required)",
	}
	rootCmd.PersistentFlags().StringVarP(&Flag.Source, args[0], shortArgs[0], "", usages[0])
	rootCmd.PersistentFlags().StringVarP(&Flag.Package, args[1], shortArgs[1], "", usages[1])

	for _, arg := range args {
		err := rootCmd.MarkPersistentFlagRequired(arg)
		if err != nil {
			return nil, fmt.Errorf("initialize command error: %s", err)
		}
	}

	// set sub commands
	rootCmd.AddCommand(initializeFocCommand())

	return rootCmd, nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	rootCmd, err := initializeRootCmd()
	if err != nil {
		return err
	}

	return rootCmd.Execute()
}

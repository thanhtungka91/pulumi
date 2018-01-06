// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package cmd

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/pulumi/pulumi/pkg/util/cmdutil"
)

func newStackExportCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "export",
		Args:  cmdutil.MaximumNArgs(0),
		Short: "Export a stack's deployment to standard out.\n",
		Long: "Export a stack's deployment to standard out.\n" +
			"\n" +
			"The deployment can then be hand-edited and used to update the stack via\n" +
			"`pulumi stack import`. This process may be used to correct inconsistencies\n" +
			"in a stack's state due to failed deployments, manual changes to cloud\n" +
			"resources, etc.",
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {
			// Fetch the current stack and export its deployment
			s, err := requireCurrentStack()
			if err != nil {
				return err
			}

			deployment, err := s.ExportDeployment()
			if err != nil {
				return err
			}

			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "    ")
			if err = enc.Encode(deployment); err != nil {
				return errors.Wrap(err, "could not export deployment")
			}
			return nil
		}),
	}
}

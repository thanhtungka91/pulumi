// Copyright 2016 Pulumi, Inc. All rights reserved.

package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func newNutVerifyCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "verify [package]",
		Short: "Check that a Nut and its NutIL are correct",
		Long: "Check that a Nut and its NutIL are correct\n" +
			"\n" +
			"A Nut contains intermediate language (NutIL) that encodes symbols,\n" +
			"definitions, and executable code.  This NutIL must obey a set of specific rules\n" +
			"for it to be valid.  Otherwise, evaluating it will fail.\n" +
			"\n" +
			"The verify command thoroughly checks the NutIL against these rules, and issues\n" +
			"errors anywhere it doesn't obey them.  This is generally useful for tools developers\n" +
			"and can ensure that Nuts do not fail at runtime, when such invariants are checked.",
		Run: runFunc(func(cmd *cobra.Command, args []string) error {
			// Create a compiler object and perform the verification.
			if !verify(cmd, args) {
				return errors.New("verification failed")
			}
			return nil
		}),
	}

	return cmd
}
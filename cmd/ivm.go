/*
Copyright © 2024 jaronnie <jaron@jaronnie.com>

*/

package cmd

import (
	"github.com/jzero-io/jzero/internal/ivm/ivminit"
	"github.com/spf13/cobra"
)

// ivmCmd represents the interface version manage command
var ivmCmd = &cobra.Command{
	Use:   "ivm",
	Short: "jzero interface version manage",
	Long:  `jzero interface version manage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var ivmInitCmd = &cobra.Command{
	Use:   "init",
	Short: "jzero ivm init",
	Long:  `jzero ivm init`,
	RunE:  ivminit.Init,
}

var ivmAddCmd = &cobra.Command{
	Use:   "add",
	Short: "jzero ivm add",
	Long:  `jzero ivm add`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(ivmCmd)

	{
		ivmCmd.AddCommand(ivmInitCmd)

		ivmInitCmd.Flags().StringVarP(&ivminit.Version, "version", "v", "v1", "jzero ivm init")
		ivmInitCmd.Flags().StringVarP(&ivminit.Style, "style", "", "gozero", "The file naming format, see [https://github.com/zeromicro/go-zero/blob/master/tools/goctl/config/readme.md]")
		ivmInitCmd.Flags().BoolVarP(&ivminit.RemoveSuffix, "remove-suffix", "", false, "remove suffix Handler and Logic on filename or file content")
	}

	{
		ivmCmd.AddCommand(ivmAddCmd)
	}
}
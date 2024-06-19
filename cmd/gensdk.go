/*
Copyright © 2024 jaronnie <jaron@jaronnie.com>

*/

package cmd

import (
	"github.com/pkg/errors"
	"path/filepath"

	"github.com/jzero-io/jzero/internal/gen/gensdk"

	"github.com/jzero-io/jzero/embeded"
	"github.com/spf13/cobra"
)

// genSdkCmd represents the gen sdk command
var genSdkCmd = &cobra.Command{
	Use:   "sdk",
	Short: "jzero gensdk",
	Long:  `jzero gensdk. Generate sdk client by api file and proto file`,
	PreRun: func(_ *cobra.Command, _ []string) {
		gensdk.Version = Version
		if gensdk.Language == "go" {
			if gensdk.Module == "" {
				cobra.CheckErr(errors.New("must specify a module to generate sdk for"))
			}
		}
	},
	RunE: gensdk.GenSdk,
}

func init() {
	genSdkCmd.Flags().StringVarP(&gensdk.Language, "language", "l", "go", "set language")
	genSdkCmd.Flags().StringVarP(&gensdk.Dir, "dir", "d", "sdk", "set dir")
	_ = genSdkCmd.MarkFlagRequired("dir")

	genSdkCmd.Flags().StringVarP(&gensdk.WorkingDir, "working-dir", "w", "", "set working dir")

	genSdkCmd.Flags().StringVarP(&gensdk.Module, "module", "m", "", "set module name")
	//_ = genSdkCmd.MarkFlagRequired("module")

	genSdkCmd.Flags().StringVarP(&gensdk.ApiDir, "api-dir", "", filepath.Join("desc", "api"), "set input api dir")
	genSdkCmd.Flags().StringVarP(&gensdk.ProtoDir, "proto-dir", "", filepath.Join("desc", "proto"), "set input proto dir")
	genSdkCmd.Flags().BoolVarP(&gensdk.WarpResponse, "warp-response", "", false, "warp response: code, data, message")
	genSdkCmd.Flags().StringVarP(&gensdk.Scope, "scope", "", "", "set scope name")
	genSdkCmd.Flags().StringVarP(&embeded.Home, "home", "", "", "set template home")
}

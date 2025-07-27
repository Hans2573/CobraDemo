package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// 版本信息
	Version   = "1.0.0"
	BuildTime = "2024-01-01"
	GitCommit = "unknown"
)

// versionCmd 表示版本命令
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示版本信息",
	Long:  `显示应用程序的版本信息，包括版本号、构建时间和Git提交哈希。`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("CobraDemo 版本信息:\n")
		fmt.Printf("  版本: %s\n", Version)
		fmt.Printf("  构建时间: %s\n", BuildTime)
		fmt.Printf("  Git提交: %s\n", GitCommit)
	},
}

// version 没有子命令
// func init() {

// }

package cmd

import (
	"fmt"

	"github.com/Hans2573/CobraDemo/cmd/config"
	"github.com/spf13/cobra"
)

// configCmd 配置管理命令
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "配置管理",
	Long:  "配置管理相关的命令",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("⚙️  配置管理系统")
		fmt.Println("================")
		fmt.Println()
		fmt.Println("🔧 可用命令:")
		fmt.Println("  show    - 显示当前配置")
		fmt.Println("  set     - 设置配置项")
		fmt.Println("  backup  - 备份管理")
		fmt.Println()
		fmt.Println("💡 使用示例:")
		fmt.Println("  cobrademo config show")
		fmt.Println("  cobrademo config set database mysql://localhost")
		fmt.Println("  cobrademo config backup list")
		fmt.Println()
		fmt.Println("📖 使用 'cobrademo config <command> --help' 获取更多帮助")
	},
}

func init() {
	// 注册基本命令
	configCmd.AddCommand(config.ShowCmd)
	configCmd.AddCommand(config.SetCmd)

	// 注册二级子命令
	configCmd.AddCommand(config.BackupCmd)
}

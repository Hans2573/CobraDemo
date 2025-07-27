package cmd

import (
	"fmt"

	"github.com/Hans2573/CobraDemo/cmd/tools"
	"github.com/spf13/cobra"
)

// toolsCmd 工具命令
var toolsCmd = &cobra.Command{
	Use:   "tools",
	Short: "工具集",
	Long:  "各种实用工具",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🛠️  实用工具集")
		fmt.Println("==============")
		fmt.Println()
		fmt.Println("🔨 可用工具:")
		fmt.Println("  hash  - 计算文本哈希值")
		fmt.Println("  time  - 显示当前时间信息")
		fmt.Println()
		fmt.Println("💡 使用示例:")
		fmt.Println("  cobrademo tools hash 'hello world'")
		fmt.Println("  cobrademo tools time")
		fmt.Println()
		fmt.Println("📖 使用 'cobrademo tools <command> --help' 获取更多帮助")
	},
}

func init() {
	// 注册基本命令
	toolsCmd.AddCommand(tools.HashCmd)
	toolsCmd.AddCommand(tools.TimeCmd)
}

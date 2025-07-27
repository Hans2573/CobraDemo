package cmd

import (
	"fmt"

	"github.com/Hans2573/CobraDemo/cmd/user"
	"github.com/spf13/cobra"
)

// userCmd 用户管理命令
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "用户管理",
	Long:  "用户管理相关的命令",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📋 用户管理系统")
		fmt.Println("================")
		fmt.Println()
		fmt.Println("🔍 可用命令:")
		fmt.Println("  list        - 列出所有用户")
		fmt.Println("  create      - 创建新用户")
		fmt.Println("  permission  - 管理用户权限")
		fmt.Println()
		fmt.Println("💡 使用示例:")
		fmt.Println("  cobrademo user list")
		fmt.Println("  cobrademo user create john")
		fmt.Println("  cobrademo user permission list 1")
		fmt.Println()
		fmt.Println("📖 使用 'cobrademo user <command> --help' 获取更多帮助")
	},
}

func init() {
	userCmd.AddCommand(user.ListCmd)
	userCmd.AddCommand(user.CreateCmd)
	userCmd.AddCommand(user.PermissionCmd)
}

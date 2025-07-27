package user

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// 命令参数变量
var (
	listFormat  string
	listActive  bool
	listLimit   int
	createEmail string
	createRole  string
	createForce bool
)

// ListCmd 列出用户
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有用户",
	Long:  "列出系统中的用户，支持多种输出格式和过滤选项",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("📋 用户列表 (格式: %s)\n", listFormat)
		fmt.Println("==================")

		users := []map[string]string{
			{"id": "1", "name": "张三", "email": "zhangsan@example.com", "role": "admin", "active": "true"},
			{"id": "2", "name": "李四", "email": "lisi@example.com", "role": "user", "active": "true"},
			{"id": "3", "name": "王五", "email": "wangwu@example.com", "role": "user", "active": "false"},
		}

		// 过滤活跃用户
		if listActive {
			fmt.Println("🟢 仅显示活跃用户")
		}

		count := 0
		for _, user := range users {
			if listActive && user["active"] != "true" {
				continue
			}
			if listLimit > 0 && count >= listLimit {
				break
			}

			switch strings.ToLower(listFormat) {
			case "json":
				fmt.Printf(`{"id":"%s","name":"%s","email":"%s","role":"%s","active":%s}%s`,
					user["id"], user["name"], user["email"], user["role"], user["active"], "\n")
			case "yaml":
				fmt.Printf("- id: %s\n  name: %s\n  email: %s\n  role: %s\n  active: %s\n",
					user["id"], user["name"], user["email"], user["role"], user["active"])
			default: // table
				status := "🔴"
				if user["active"] == "true" {
					status = "🟢"
				}
				fmt.Printf("%s ID:%s | %s (%s) | %s\n", status, user["id"], user["name"], user["role"], user["email"])
			}
			count++
		}

		if listLimit > 0 {
			fmt.Printf("\n📊 显示了前 %d 个用户\n", count)
		}
	},
}

// CreateCmd 创建用户
var CreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "创建用户",
	Long:  "创建新用户，支持设置邮箱、角色等属性",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		fmt.Printf("👤 创建用户: %s\n", name)
		fmt.Println("=================")

		if createEmail != "" {
			fmt.Printf("📧 邮箱: %s\n", createEmail)
		}

		if createRole != "" {
			fmt.Printf("🎭 角色: %s\n", createRole)
		} else {
			fmt.Printf("🎭 角色: user (默认)\n")
		}

		if createForce {
			fmt.Println("⚠️  使用强制模式，将覆盖已存在的用户")
		}

		fmt.Println("✅ 用户创建成功!")
	},
}

func init() {
	// user list 命令参数
	ListCmd.Flags().StringVarP(&listFormat, "format", "f", "table", "输出格式 (table|json|yaml)")
	ListCmd.Flags().BoolVarP(&listActive, "active", "a", false, "仅显示活跃用户")
	ListCmd.Flags().IntVarP(&listLimit, "limit", "l", 0, "限制显示数量 (0=不限制)")

	// user create 命令参数
	CreateCmd.Flags().StringVarP(&createEmail, "email", "e", "", "用户邮箱地址")
	CreateCmd.Flags().StringVarP(&createRole, "role", "r", "", "用户角色 (admin|user|guest)")
	CreateCmd.Flags().BoolVarP(&createForce, "force", "", false, "强制创建，覆盖已存在用户")
}

package config

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// 命令参数变量
var (
	showFormat  string
	showSection string
	setType     string
	setValidate bool
	setBackup   bool
)

// ShowCmd 显示配置
var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "显示配置",
	Long:  "显示系统配置信息，支持多种输出格式和分节显示",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("⚙️  系统配置 (格式: %s)\n", showFormat)
		fmt.Println("====================")

		configs := map[string]map[string]string{
			"database": {
				"host":     "localhost",
				"port":     "3306",
				"name":     "cobrademo",
				"user":     "admin",
				"password": "******",
			},
			"server": {
				"host": "0.0.0.0",
				"port": "8080",
				"ssl":  "false",
			},
			"logging": {
				"level":  "info",
				"file":   "/var/log/cobrademo.log",
				"rotate": "true",
			},
		}

		// 过滤特定配置节
		sectionsToShow := make(map[string]map[string]string)
		if showSection != "" {
			if section, exists := configs[showSection]; exists {
				sectionsToShow[showSection] = section
				fmt.Printf("📂 仅显示配置节: %s\n", showSection)
			} else {
				fmt.Printf("❌ 配置节 '%s' 不存在\n", showSection)
				return
			}
		} else {
			sectionsToShow = configs
		}

		// 根据格式输出
		switch strings.ToLower(showFormat) {
		case "json":
			fmt.Println("{")
			sectionCount := 0
			for sectionName, section := range sectionsToShow {
				fmt.Printf(`  "%s": {`, sectionName)
				itemCount := 0
				for key, value := range section {
					if itemCount > 0 {
						fmt.Print(",")
					}
					fmt.Printf(`"%s": "%s"`, key, value)
					itemCount++
				}
				fmt.Print("}")
				if sectionCount < len(sectionsToShow)-1 {
					fmt.Print(",")
				}
				fmt.Println()
				sectionCount++
			}
			fmt.Println("}")
		case "yaml":
			for sectionName, section := range sectionsToShow {
				fmt.Printf("%s:\n", sectionName)
				for key, value := range section {
					fmt.Printf("  %s: %s\n", key, value)
				}
			}
		default: // table
			for sectionName, section := range sectionsToShow {
				fmt.Printf("\n📁 [%s]\n", strings.ToUpper(sectionName))
				for key, value := range section {
					fmt.Printf("  %s = %s\n", key, value)
				}
			}
		}
	},
}

// SetCmd 设置配置
var SetCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "设置配置",
	Long:  "设置系统配置项，支持类型验证和备份功能",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		fmt.Printf("⚙️  设置配置项\n")
		fmt.Println("===============")

		if setBackup {
			fmt.Println("💾 正在创建配置备份...")
			fmt.Println("✅ 备份创建成功: config_backup_20240127.yaml")
		}

		if setValidate {
			fmt.Printf("🔍 正在验证配置值...")
			// 模拟验证逻辑
			if strings.Contains(key, "port") && setType == "int" {
				fmt.Println(" ✅ 端口号验证通过")
			} else if strings.Contains(key, "host") && setType == "string" {
				fmt.Println(" ✅ 主机地址验证通过")
			} else {
				fmt.Println(" ✅ 基本格式验证通过")
			}
		}

		typeInfo := ""
		if setType != "" {
			typeInfo = fmt.Sprintf(" (类型: %s)", setType)
		}

		fmt.Printf("📝 设置: %s = %s%s\n", key, value, typeInfo)
		fmt.Println("✅ 配置设置成功!")

		if setValidate || setBackup {
			fmt.Println("\n💡 提示: 配置更改已生效，重启服务以应用所有更改")
		}
	},
}

func init() {
	// config show 命令参数
	ShowCmd.Flags().StringVarP(&showFormat, "format", "f", "table", "输出格式 (table|json|yaml)")
	ShowCmd.Flags().StringVarP(&showSection, "section", "s", "", "仅显示指定配置节 (database|server|logging)")

	// config set 命令参数
	SetCmd.Flags().StringVarP(&setType, "type", "t", "", "值类型 (string|int|bool|float)")
	SetCmd.Flags().BoolVarP(&setValidate, "validate", "v", false, "验证配置值格式")
	SetCmd.Flags().BoolVarP(&setBackup, "backup", "b", false, "设置前创建备份")
}

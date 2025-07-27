package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// 用于存储配置文件的标志
	cfgFile string
	// 用于存储日志级别的标志
	logLevel string
)

// rootCmd 表示没有调用子命令时的基础命令
var rootCmd = &cobra.Command{
	Use:   "cobrademo",
	Short: "一个使用Cobra构建的多级命令行工具演示",
	Long: `这是一个使用Cobra包构建的多级命令行工具演示程序。
	
它展示了如何创建具有子命令、标志和参数的命令行应用程序。
	
示例用法:
  cobrademo version          # 显示版本信息
  cobrademo config show      # 显示配置信息
  cobrademo user list        # 列出用户
  cobrademo user create      # 创建用户`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World! This is the root command.")
	},
}

// Execute 添加所有子命令到根命令并设置标志
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// 在命令执行之前，先运行 initConfig 函数。
	cobra.OnInitialize(initConfig)

	// 全局标志，可用于所有子命令
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "配置文件路径 (默认为 $HOME/.cobrademo.yaml)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "日志级别 (debug|info|warn|error)")

	// 本地标志，仅用于根命令
	rootCmd.Flags().BoolP("toggle", "t", false, "帮助消息的切换")

	// 绑定标志到viper
	viper.BindPFlag("log-level", rootCmd.PersistentFlags().Lookup("log-level"))

	// 注册所有子命令
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(userCmd)
	rootCmd.AddCommand(toolsCmd)
}

// initConfig 读取配置文件和环境变量
func initConfig() {
	if cfgFile != "" {
		// 使用指定的配置文件
		viper.SetConfigFile(cfgFile)
	} else {
		// 搜索配置文件
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// 搜索配置文件名称
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobrademo")
	}

	// 读取环境变量
	viper.AutomaticEnv()

	// 如果找到配置文件，则读取它
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "使用配置文件:", viper.ConfigFileUsed())
	}
}

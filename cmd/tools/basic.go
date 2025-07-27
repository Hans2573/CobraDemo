package tools

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// 命令参数变量
var (
	hashAlgorithm string
	hashUppercase bool
	hashFromFile  string
	timeFormat    string
	timeZone      string
	timeUnixOnly  bool
)

// HashCmd 哈希工具
var HashCmd = &cobra.Command{
	Use:   "hash [text]",
	Short: "计算哈希值",
	Long:  "计算文本或文件的哈希值，支持多种哈希算法",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var input string
		var isFromFile bool

		// 确定输入源
		if hashFromFile != "" {
			// 从文件读取
			content, err := os.ReadFile(hashFromFile)
			if err != nil {
				fmt.Printf("❌ 读取文件失败: %v\n", err)
				return
			}
			input = string(content)
			isFromFile = true
			fmt.Printf("📁 从文件读取: %s\n", hashFromFile)
		} else if len(args) > 0 {
			// 从参数读取
			input = args[0]
		} else {
			fmt.Println("❌ 请提供要计算哈希的文本或使用 --file 参数指定文件")
			return
		}

		fmt.Printf("🔐 计算哈希值 (算法: %s)\n", strings.ToUpper(hashAlgorithm))
		fmt.Println("========================")

		if !isFromFile {
			fmt.Printf("📝 输入文本: '%s'\n", input)
		}
		fmt.Println()

		// 选择哈希算法
		var hasher hash.Hash
		switch strings.ToLower(hashAlgorithm) {
		case "md5":
			hasher = md5.New()
		case "sha1":
			hasher = sha1.New()
		case "sha256":
			hasher = sha256.New()
		case "sha512":
			hasher = sha512.New()
		default:
			// 默认计算所有常用算法
			fmt.Println("🔍 计算多种哈希算法:")
			algorithms := []string{"md5", "sha1", "sha256", "sha512"}
			for _, alg := range algorithms {
				result := calculateHash(input, alg, hashUppercase)
				fmt.Printf("  %s: %s\n", strings.ToUpper(alg), result)
			}
			return
		}

		// 计算指定算法
		hasher.Write([]byte(input))
		result := fmt.Sprintf("%x", hasher.Sum(nil))

		if hashUppercase {
			result = strings.ToUpper(result)
		}

		fmt.Printf("✅ %s: %s\n", strings.ToUpper(hashAlgorithm), result)
	},
}

// TimeCmd 时间工具
var TimeCmd = &cobra.Command{
	Use:   "time",
	Short: "显示时间",
	Long:  "显示当前时间信息，支持自定义格式和时区",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()

		// 处理时区
		if timeZone != "" {
			loc, err := time.LoadLocation(timeZone)
			if err != nil {
				fmt.Printf("❌ 无效的时区: %s\n", timeZone)
				fmt.Println("💡 常用时区: UTC, Asia/Shanghai, America/New_York, Europe/London")
				return
			}
			now = now.In(loc)
			fmt.Printf("🌍 时区: %s\n", timeZone)
		}

		if timeUnixOnly {
			// 仅显示Unix时间戳
			fmt.Printf("%d\n", now.Unix())
			return
		}

		fmt.Println("🕐 时间信息")
		fmt.Println("===========")

		// 根据格式显示时间
		if timeFormat != "" {
			fmt.Printf("📅 自定义格式: %s\n", now.Format(timeFormat))
		} else {
			// 显示多种格式
			fmt.Printf("📅 标准格式: %s\n", now.Format("2006-01-02 15:04:05"))
			fmt.Printf("📅 RFC3339:  %s\n", now.Format(time.RFC3339))
			fmt.Printf("📅 ISO8601:  %s\n", now.Format("2006-01-02T15:04:05Z07:00"))
		}

		fmt.Printf("⏰ Unix时间戳: %d\n", now.Unix())
		fmt.Printf("⏰ 毫秒时间戳: %d\n", now.UnixMilli())
		fmt.Printf("🌅 星期: %s\n", now.Weekday().String())
		fmt.Printf("📆 年中第%d天\n", now.YearDay())
	},
}

// 辅助函数：计算哈希值
func calculateHash(input, algorithm string, uppercase bool) string {
	var hasher hash.Hash
	switch strings.ToLower(algorithm) {
	case "md5":
		hasher = md5.New()
	case "sha1":
		hasher = sha1.New()
	case "sha256":
		hasher = sha256.New()
	case "sha512":
		hasher = sha512.New()
	default:
		return "unsupported algorithm"
	}

	hasher.Write([]byte(input))
	result := fmt.Sprintf("%x", hasher.Sum(nil))

	if uppercase {
		return strings.ToUpper(result)
	}
	return result
}

func init() {
	// tools hash 命令参数
	HashCmd.Flags().StringVarP(&hashAlgorithm, "algorithm", "a", "", "哈希算法 (md5|sha1|sha256|sha512)，留空显示所有")
	HashCmd.Flags().BoolVarP(&hashUppercase, "uppercase", "u", false, "输出大写哈希值")
	HashCmd.Flags().StringVarP(&hashFromFile, "file", "f", "", "从文件读取内容计算哈希")

	// tools time 命令参数
	TimeCmd.Flags().StringVarP(&timeFormat, "format", "f", "", "自定义时间格式 (Go时间格式)")
	TimeCmd.Flags().StringVarP(&timeZone, "timezone", "z", "", "指定时区 (如: UTC, Asia/Shanghai)")
	TimeCmd.Flags().BoolVarP(&timeUnixOnly, "unix", "u", false, "仅显示Unix时间戳")
}

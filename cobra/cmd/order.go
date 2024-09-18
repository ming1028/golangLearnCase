package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "命令名称",
	Short: "简短介绍",
	Long:  "详细介绍",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("命令具体操作", cmd.Use, len(args))
		flag1, _ := cmd.Flags().GetString("flag1")
		fmt.Println(flag1)
		flag2, _ := cmd.Flags().GetString("flag2")
		fmt.Println(flag2)
	},
}

func Execute() {
	rootCmd.PersistentFlags().String("flag1", "1", "flag1参数使用用途")
	rootCmd.Flags().String("flag2", "2", "flag2参数使用用途")
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

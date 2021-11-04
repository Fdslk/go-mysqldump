package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dianpuId string

func init() {
	initLocalShop.Flags().StringVarP(&dianpuId, "dianpuId", "d", "", "需要导出数据的店铺Id")
	initLocalShop.MarkFlagRequired("dianpuId")

	rootCmd.AddCommand(initLocalShop)
}

var initLocalShop = &cobra.Command{
	Use:   "localize",
	Short: "店铺数据本地化",
	Long:  "通过店铺id导出当前店铺的本地资源数据",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(dianpuId)
	},
}

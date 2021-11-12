package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	common "github.com/dumpsql/Common"
	config "github.com/dumpsql/Config"
	"github.com/spf13/cobra"
)

var (
	dbname = "localEmenu"
	newMap map[string]interface{}
)

func init() {
	versionCmd.Flags().StringVarP(&dianpuId, "dianpuId", "d", "", "商店id的参数")
	versionCmd.MarkFlagRequired("dianpuId")

	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "shopInfo",
	Short: "获取店铺信息",
	Long:  `通过店铺id获取店铺相关信息`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getShopInfo())
	},
}

func getShopInfo() (string, error) {
	dbConfig, _ := config.NewConfig("Config/config.yaml")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbConfig.DataBase.UserName,
		dbConfig.DataBase.PassWord,
		dbConfig.DataBase.HostName,
		strconv.Itoa(dbConfig.DataBase.Port), dbname))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var shopObject shop

	sql := fmt.Sprintf("Select MingCheng, FuZeRenXingMing, GuDingDianHua, DiZhi, IsYingYe from tb_dianpu where dianpuid = %s and yuyan = 1", dianpuId)

	err = db.QueryRow(sql).Scan(&shopObject.MingCheng,
		&shopObject.FuZeRenXingMing,
		&shopObject.GuDingDianHua,
		&shopObject.DiZhi,
		&shopObject.IsYingYe)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return common.GetColorfulString(shopObject)
}

type shop struct {
	MingCheng, FuZeRenXingMing, GuDingDianHua, DiZhi string
	IsYingYe                                         int
}

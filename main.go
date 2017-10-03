package main

import (
	"flag"
	"fmt"
	"github.com/iikira/baidu-tools/tieba"
)

var (
	bduss     = flag.String("b", "", "Baidu BDUSS.")
	tiebaName = flag.String("t", "", "Tieba name.")
	user      = flag.String("u", "", "Tieba username.")
	day       = flag.Int("d", 1, "Time(day) for ban, only supports 1, 3, 10.")
	reason    = flag.String("r", "null", "Reason for ban.")
)

type ban struct {
	*tieba.Tieba
	preBanUser, // 准备封禁百度用户名
	reason string // 封禁理由
	day int // 封禁天数
	bar tieba.Bar
}

func main() {
	fmt.Printf("Tieba-Ban-User %s, Last Update: 2017-09-26, \nGithub: https://github.com/iikira/Tieba-Ban-User\n\n", version)

	//解析flag参数
	flag.Parse()

	//判断是否输入数据
	if *bduss == "" || *tiebaName == "" || *user == "" {
		print("Not enough input data. \n Try -h for more infomation.\n")
		return
	}

	if *day != 1 && *day != 3 && *day != 10 {
		print("Ban day invalid. \n Try -h for more infomation.\n")
		return
	}

	//初始化数据，执行封禁
	ban := ban{
		preBanUser: *user,
		bar:        tieba.Bar{Name: *tiebaName},
		reason:     *reason,
		day:        *day,
	}

	ban.bar.Fid, _ = tieba.GetTiebaFid(ban.bar.Name) //初始化数据，获取贴吧fid

	ban.Tieba, _ = tieba.NewWithBDUSS(*bduss)
	ban.banUser()
}

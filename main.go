package main

import (
	"flag"
	"fmt"
	"github.com/iikira/Tieba-Cloud-Sign-Backend/baiduUtil"
)

var ( // -ldflags "-X main.version=1.0"
	version = "1.0"

	bduss  = flag.String("b", "", "Baidu BDUSS.")
	tieba  = flag.String("t", "", "Tieba name.")
	user   = flag.String("u", "", "Tieba username.")
	day    = flag.Int("d", 1, "Time(day) for ban, only supports 1, 3, 10.")
	reason = flag.String("r", "null", "Reason for ban.")
)

type ban struct {
	BDUSS, // 百度BDUSS
	user, // 百度用户名
	reason string // 封禁理由
	day int // 封禁天数
	ba
}

func main() {
	fmt.Printf("Tieba-Ban-User v%s, Last Update: 2017-07-29, \nGithub: https://github.com/iikira/Tieba-Ban-User\n\n", version)

	//解析flag参数
	flag.Parse()

	//判断是否输入数据
	if *bduss == "" || *tieba == "" || *user == "" {
		print("Not enough input data. \n Try -h for more infomation.\n")
		return
	}

	if *day != 1 && *day != 3 && *day != 10 {
		print("Ban day invalid. \n Try -h for more infomation.\n")
		return
	}

	//初始化数据，获取贴吧fid
	ba := ba{
		name: *tieba,
	}
	ba.fid, _ = baiduUtil.GetTiebaFid(ba.name)

	//初始化数据，执行封禁
	ban := ban{
		BDUSS:  *bduss,
		user:   *user,
		ba:     ba,
		reason: *reason,
		day:    *day,
	}
	ban.banUser()
}

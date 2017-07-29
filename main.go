package main

import "flag"
import "fmt"

var (
	bduss  = flag.String("b", "", "Baidu BDUSS.")
	tieba  = flag.String("t", "", "Tieba name.")
	user   = flag.String("u", "", "Tieba username.")
	day    = flag.Int("d", 1, "Time(day) for ban, only supports 1, 3, 10.")
	reason = flag.String("r", "null", "Reason for ban.")
)

type ban struct {
	BDUSS, user, reason string
	day                 int
	ba
}

func main() {
	fmt.Println("Tieba-Ban-User v1.0, Last Update: 2017-07-29")
	fmt.Println("Author: iikira")
	fmt.Println()

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
		tieba: *tieba,
	}
	ba.getFid()

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

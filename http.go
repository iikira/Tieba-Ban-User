package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"my/Baidu-Tools/util"
)

<<<<<<< HEAD
// banUser 执行 封禁
=======
type ba struct {
	fid  string
	name string
}

// getTbs 获取贴吧tbs
func getTbs(BDUSS string) string {
	header := map[string]string{
		"Cookie": "BDUSS=" + BDUSS,
	}
	body, err := baiduUtil.Fetch("GET", "http://tieba.baidu.com/dc/common/tbs", nil, nil, header)
	if err != nil {
		return ""
	}
	json, _ := simplejson.NewJson(body)
	return json.Get("tbs").MustString()
}

// banUser 执行封禁
>>>>>>> 27eaacfc803494055859b6214cec2fbb1d94b40b
func (b *ban) banUser() {
	post := map[string]string{
		"BDUSS":  b.BDUSS,
		"day":    fmt.Sprintf("%d", b.day),
		"fid":    b.bar.Fid,
		"ntn":    "banid",
		"reason": b.reason,
		"tbs":    b.Tbs,
		"un":     b.preBanUser,
		"word":   b.bar.Name,
		"z":      "1111111111",
	}

	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"Cookie":       "ka=open",
		"net":          "1",
		"User-Agent":   "bdtb for Android 6.9.2.1",
		"Connection":   "Keep-Alive",
	}
	baiduUtil.TiebaClientSignature(post)
	body, err := baiduUtil.Fetch("POST", "http://tieba.baidu.com/c/c/bawu/commitprison", nil, post, header)

	baiduUtil.PrintErrAndExit("执行封禁时网络错误：", err)

	banResult, err := simplejson.NewJson(body)
	baiduUtil.PrintErrAndExit("解析Json数据失败：", err)

	errorCode := banResult.Get("error_code").MustString()
	errorMsg := banResult.Get("error_msg").MustString("未找到错误原因, 可能已成功封禁, 源响应body: " + string(body))

	fmt.Printf("代码: %s, 信息: %s\n", errorCode, errorMsg)
}

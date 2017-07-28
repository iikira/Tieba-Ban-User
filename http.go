package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/iikira/Tieba-Cloud-Sign-Backend/baiduUtil"
	"regexp"
)

type ba struct {
	fid   string
	tieba string
}

// getFid 获取贴吧fid
func (b *ba) getFid() {
	body, _, _ := baiduUtil.Fetch("http://tieba.baidu.com/mo/m?kw="+b.tieba, nil, nil)
	rawFid := regexp.MustCompile(`<input type="hidden" name="fid" value="(.*?)"/>`).FindStringSubmatch(body)
	if len(rawFid) <= 1 {
		b.fid = "0"
	} else {
		b.fid = rawFid[1]
	}
}

// getTbs 获取贴吧tbs
func getTbs(BDUSS string) string {
	Data := map[string]string{
		"Cookie": "BDUSS=" + BDUSS,
	}
	body, _, err := baiduUtil.Fetch("http://tieba.baidu.com/dc/common/tbs", nil, &Data)
	if err != nil {
		return ""
	}
	json, _ := simplejson.NewJson([]byte(body))
	return json.Get("tbs").MustString()
}

// banUser 执行封禁
func (b *ban) banUser() {
	postData := map[string]string{
		"BDUSS":  b.BDUSS,
		"day":    fmt.Sprintf("%d", b.day),
		"fid":    b.fid,
		"ntn":    "banid",
		"reason": b.reason,
		"tbs":    getTbs(b.BDUSS),
		"un":     b.user,
		"word":   b.tieba,
		"z":      "1111111111",
	}

	tiebaHeaderData := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"Cookie":       "ka=open",
		"net":          "1",
		"User-Agent":   "bdtb for Android 6.9.2.1",
		"Connection":   "Keep-Alive",
	}
	postData["sign"] = baiduUtil.ClientSignature(&postData)
	body, _, err := baiduUtil.Fetch("http://tieba.baidu.com/c/c/bawu/commitprison", &postData, &tiebaHeaderData)

	baiduUtil.PrintErrAndExit("执行封禁时网络错误：", err)

	banResult, err := simplejson.NewJson([]byte(body))
	baiduUtil.PrintErrAndExit("解析Json数据失败：", err)

	errorCode := banResult.Get("error_code").MustString()
	errorMsg := banResult.Get("error_msg").MustString("未找到错误原因, 源响应body: " + body)

	fmt.Printf("代码: %s, 信息: %s\n", errorCode, errorMsg)
}

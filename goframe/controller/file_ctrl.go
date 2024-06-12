package controller

import (
	"io/ioutil"
	"os"

	"github.com/gogf/gf/v2/net/ghttp"
)

type fileCtrl struct{}

var FileCtrl fileCtrl

// Get 服务健康检查
func (ctrl *fileCtrl) Get(r *ghttp.Request) {
	file, err := os.Open("D:\\development\\goProject\\study\\httpServerStudy\\goframe\\controller\\levelData.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 读取文件内容
	body, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	r.Response.WriteJson(body)
}

package controller

import (
	"io/ioutil"
	"os"

	"github.com/gogf/gf/v2/net/ghttp"
)

type fileCtrl struct{}

var FileCtrl fileCtrl

// GetLevelConfig 获取关卡配置
func (ctrl *fileCtrl) GetLevelConfig(r *ghttp.Request) {
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

func (ctrl *fileCtrl) GetImage(r *ghttp.Request) {
	file, err := os.Open("D:\\development\\goProject\\study\\httpServerStudy\\goframe\\controller\\img.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//// 读取文件内容
	body, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	//// 创建一个新的图片文件
	//newFilePath := "D:\\development\\goProject\\study\\httpServerStudy\\goframe\\controller\\new_img.png"
	//newFile, err := os.Create(newFilePath)
	//if err != nil {
	//	log.Fatalf("Failed to create new file: %v", err)
	//}
	//defer newFile.Close()
	//
	//// 将读取的内容写入新文件
	//_, err = newFile.Write(body)
	//if err != nil {
	//	log.Fatalf("Failed to write to new file: %v", err)
	//}

	r.Response.WriteJson(body)
}

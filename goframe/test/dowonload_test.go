package test

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func Test_Download(t *testing.T) {
	dir, _ := os.Getwd()
	fmt.Println("dir-->", dir)
	err := downloadFile("https://89t-pub-file.s3.us-west-2.amazonaws.com/nuts/61e3cd98d38849f8386665ff5f5bfbb4", "levelData.json")
	if err != nil {
		panic(err)
	}
}

func downloadFile(url string, fileName string) error {
	// 创建HTTP GET请求
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// 创建文件
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将响应的Body数据写入文件
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

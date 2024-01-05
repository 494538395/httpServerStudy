package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//url := "http://127.0.0.1:9997/openSearch"
	//url := "http://127.0.0.1:9997/redis"
	url := "http://127.0.0.1:9997/mysql"

	client := &http.Client{}

	// 开辟协程数
	for i := 0; i < 2; i++ {
		go func() {
			// 每个协程请求数
			for k := 0; k < 5; k++ {
				go func() {
					req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
					if err != nil {
						fmt.Printf("无法创建请求：%v\n", err)
						return
					}
					req.Header.Set("Content-Type", "application/json")

					resp, err := client.Do(req)
					if err != nil {
						fmt.Printf("请求发生错误：%v\n", err)
						return
					}

					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						fmt.Printf("无法读取响应体：%v\n", err)
						resp.Body.Close()
						return
					}
					fmt.Println(string(body))

					resp.Body.Close()
				}()
				//time.Sleep(10 * time.Millisecond)
			}
		}()
	}
	select {}
}

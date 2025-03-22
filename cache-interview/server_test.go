package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 测试获取用户处理函数
func TestHandlerGetUser(t *testing.T) {
	// 创建一个新的服务器
	s := NewServer()
	// 创建一个新的HTTP服务器，使用s.handlerGetUser作为处理函数
	ts := httptest.NewServer(http.HandlerFunc(s.handlerGetUser))

	// 循环1000次
	for i := 0; i < 1000; i++ {
		// 计算id
		id := i%100 + 1
		// 发送GET请求
		rep, err := http.Get(fmt.Sprintf("%s/?id=%d", ts.URL, id))
		// 如果发生错误，输出错误信息
		if err != nil {
			t.Error(err)
		}

		// 创建一个新的User结构体
		user := &User{}
		// 解码响应体
		err = json.NewDecoder(rep.Body).Decode(user)
		// 如果发生错误，输出错误信息
		if err != nil {
			t.Error(err)
		}

		// 打印用户信息
		fmt.Printf("%+v\n", user)

	}

	// 打印数据库访问次数
	fmt.Println("db hits: ", s.hits)

}

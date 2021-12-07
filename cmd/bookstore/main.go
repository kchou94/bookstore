package main

import (
	"bookstore/server"
	"bookstore/store/factory"
	"log"
)

func main() {
	s, err := factory.New("mem") // 创建图书数据存储模块实例
	if err != nil {
		panic(err)
	}

	srv := server.NewBookStoreServer(":8080", s) // 创建http服务实例

	errChan, err := srv.ListenAndServe() // 启动http服务
	if err != nil {
		log.Println("web server start failed:", err)
		return
	}

}

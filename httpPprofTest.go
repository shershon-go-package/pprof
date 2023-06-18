package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "net/http/pprof" // 导入pprof
	"strings"
	"time"
)

func init() {
	//开启http端口,用协程的方式监听，否则会阻塞
	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			fmt.Println("pprof err:", err)
		}
	}()
}

func main2() {
	engine := gin.Default()
	engine.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "success",
		})
	})
	testPprofHeap()
	_ = engine.Run(":8080")
}

// 模拟内存使用增加
func testPprofHeap() {
	go func() {
		var stringSlice []string
		for {
			time.Sleep(time.Second * 2)
			repeat := strings.Repeat("hello,world", 50000)
			stringSlice = append(stringSlice, repeat)
			fmt.Printf("time:%d length:%d \n", time.Now().Unix(), len(stringSlice))
		}
	}()
}

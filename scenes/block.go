package scenes

import "time"

var reqCh = make(chan string)
var sqlCh = make(chan []string)
var bCh = make(chan bool)

func UseBlockDemo() {
	for i := 0; i < 1000; i++ {
		handle()
	}
}

func handle()  {
	// 等待文件写入
	go func() {
		time.Sleep(time.Second)
		bCh <- true
	}()
	<-bCh
	// 发起网络请求
	go func() {
		time.Sleep(time.Millisecond * 500)
		reqCh <- "req success"
	}()
	// 阻塞等待
	<-reqCh
	// 执行SQL查询
	go func() {
		time.Sleep(time.Millisecond * 200)
		sqlCh <- []string{"hello","world"}
	}()
	<-sqlCh
}


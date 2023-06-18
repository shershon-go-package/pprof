package tests

import (
	"strings"
	"testing"
	"time"
)

func TestWithPProf(t *testing.T) {
	ch := make(chan bool)
	go func() {
		var stringSlice []string
		for i := 0; i < 20; i++ {
			repeat := strings.Repeat("hello,world", 50000)
			stringSlice = append(stringSlice, repeat)
			time.Sleep(time.Millisecond * 500)
		}
		ch <- true
	}()
	<-ch
}

/*func TestTrace(t *testing.T) {
	// 创建文件
	file, _ := os.OpenFile("./trace.out", os.O_APPEND|os.O_CREATE, os.ModePerm)
	// 开启
	if err := trace.Start(file); err != nil {
		fmt.Println("trace start err ", err)
		return
	}
	// 关闭
	defer trace.Stop()
}*/

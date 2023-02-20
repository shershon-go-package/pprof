package tests

import (
	"fmt"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"strings"
	"testing"
	"time"
)

func TestRuntimePProf(t *testing.T) {
	// 打开文件
	f, err := os.Create("./out.pprof")
	if err != nil {
		t.Errorf("文件打开失败:%v", err)
		return
	}
	// 调用
	err = pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	if err != nil {
		t.Errorf("StartCPUProfile:%v", err)
	}
	// 测试单独函数
	testPprof()
}

// 模拟内存使用增加
func testPprof() {
	ch := make(chan bool)
	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(time.Millisecond * 200)
		}
		ch <- true
	}()
	<-ch
}

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

func TestTrace(t *testing.T) {
	// 创建文件
	file, _ := os.OpenFile("./trace.out", os.O_APPEND|os.O_CREATE, os.ModePerm)
	// 开启
	if err := trace.Start(file); err != nil {
		fmt.Println("trace start err ", err)
		return
	}
	// 关闭
	defer trace.Stop()
}

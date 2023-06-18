package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // 导入pprof
	"runtime"
	"shershon1991/go-pprof/scenes"
)

func init() {
	// 开启http端口,用协程的方式监听，否则会阻塞
	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			fmt.Println("pprof err:", err)
		}
	}()
	//默认不开启锁的竞争分析，需要调用runtime.SetMutexProfileFraction进行设置。
	runtime.SetMutexProfileFraction(1)
	// 开启block分析
	runtime.SetBlockProfileRate(1)
}

func main() {
	fmt.Println("运行测试...")
	ch := make(chan bool)
	// 堆内存
	//scenes.UseHeapDemo()
	// CPU
	//scenes.UseCpuDemo()
	// Goroutine
	//scenes.UseGoroutineDemo()
	// 锁
	//scenes.UseMutexDemo()
	// block
	scenes.UseBlockDemo()
	<-ch
}

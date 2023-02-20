package scenes

import (
	"sync"
	"time"
)

type student struct {
	idList []int
	mutex  sync.Mutex
}
func UseMutexDemo() {
	var stu student
	for i := 1 ; i < 10000; i++{
		id := i
		go func() {
			// 加锁
			stu.mutex.Lock()
			defer stu.mutex.Unlock()
			// 添加数据
			stu.idList = append(stu.idList, id)
			time.Sleep(time.Millisecond * 200)
		}()
	}
}

package scenes

import "time"

func UseGoroutineDemo()  {
	errorUse()
	trueUse()
}

func errorUse()  {
	for  {
		time.Sleep(time.Second)
		go func() {
			time.Sleep(time.Second * 30)
		}()
	}
}
func trueUse()  {
	res := make(chan []int)
	go func(ch chan <- []int) {
		var tmp []int
		for i := 0; i < 1000; i++ {
			tmp = append(tmp,i)
		}
		ch <- tmp
	}(res)
	<-res
}
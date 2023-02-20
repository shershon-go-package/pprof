package tests

import (
	"strings"
	"testing"
	"time"
)

func BenchmarkRun(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ch := make(chan bool)
		go func() {
			var stringSlice []string
			for i := 0; i < 5; i++ {
				repeat := strings.Repeat("hello,world", 50000)
				stringSlice = append(stringSlice, repeat)
				time.Sleep(time.Millisecond * 100)
			}
			ch <- true
		}()
		<-ch
	}
}

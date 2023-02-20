package scenes

import "fmt"

func UseCpuDemo()  {
	compute()
}

func compute()  {
	var total int64
	for i := 0; i < 100000000000; i++ {
		total = total + int64(i)
	}
	fmt.Println("遍历完成")
}
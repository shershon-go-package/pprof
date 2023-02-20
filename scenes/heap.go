package scenes

import (
	"fmt"
	"time"
)

func UseHeapDemo()  {
	go func() {
		var result []string
		i := 1
		for {
			body := getBody()
			girl := getGirl()
			result = append(result,body...)
			result = append(result,girl...)
			fmt.Printf("第%d轮，count:%v \n",i,len(result))
			time.Sleep(time.Second * 5)
			i++
		}
	}()
}

func getBody() []string {
	var body []string
	i := 0
	for  {
		if i > 10000 {
			break
		}
		body =append(body,fmt.Sprintf("男生%d",i))
		i++
	}
	return body
}

func getGirl() []string {
	var girl []string
	i := 0
	for  {
		if i > 10000 {
			break
		}
		girl = append(girl,fmt.Sprintf("女生%d",i))
		i++
	}
	return girl
}





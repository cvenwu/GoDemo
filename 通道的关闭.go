package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// more表示通道是否关闭
			j, more := <-jobs
			fmt.Println("j = ", j, "more = ", more, "")
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 0; j < 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done

	// 发现虽然close，但是通道中如果有还没有读取的数据，ok就会为true
	fmt.Println("-----------")
	ch := make(chan int, 10)
	ch <- 10
	ch <- 9
	ch <- 8
	_, ok := <-ch
	fmt.Println(ok)
	// 关闭通道
	close(ch)
	_, ok = <-ch
	fmt.Println(ok)
	_, ok = <-ch
	fmt.Println(ok)
	_, ok = <-ch
	fmt.Println(ok)
}

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// fmt.Print("hello world\n")
	// ch1 := make(chan int)
	// ch2 := make(chan int)
	// stopCh := make(chan struct{})
	// wg.Add(1)
	// go work(ch1, ch2, stopCh)
	// ch1 <- 1
	// ch2 <- 1
	// time.Sleep(1000 * time.Millisecond)
	// // stopCh <- struct{}{}
	// var a int

	// wg.Wait()
	// fmt.Scan(&a)
	var a = [2]byte{'[', ']'}
	fmt.Printf("%d", a[0])
}
func work(ch1, ch2 chan int, stopCh <-chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-stopCh:
			return
		case job1 := <-ch1:
			fmt.Println("receive ch1 on first loop:", job1)
		case job2 := <-ch2:
		priority:
			for {
				select {
				case job1 := <-ch1:
					fmt.Println("receive ch1 on second loop:", job1)
				default:
					time.Sleep(1000 * time.Millisecond)
					break priority
				}
			}
			// 执行job1的内容
			fmt.Println("receive ch2:", job2)
		}
	}
}

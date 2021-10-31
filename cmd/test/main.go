package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

func main() {
	//list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	products := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	productChan := list.New()
	for _, p := range products {
		productChan.PushFront(p)
	}
	var aIndex int = 0
	var bIndex int = 0
	//var wait sync.WaitGroup
	wait := make(chan int, 1)
	var end sync.WaitGroup
	end.Add(1)
	var lock sync.Mutex
	go func() {
		for {
			lock.Lock()
			p := productChan.Back()
			time.Sleep(1 * time.Second)
			if p == nil {
				fmt.Printf("A完成所有产品的运输！\n")
				return
			}
			productChan.Remove(p)
			lock.Unlock()
			for {
				if aIndex >= bIndex {
					<-wait
				} else {
					fmt.Printf("A将产品【%s】放入第【%d】个盒子！\n", p.Value, aIndex)
					aIndex += 1
					if aIndex == 9 {
						fmt.Printf("A对产品【%s】完成整个流程的运输！\n", p.Value)
						aIndex = 0
						break
					}
				}
			}
		}
	}()

	go func() {
		for {
			lock.Lock()
			p := productChan.Back()
			time.Sleep(1 * time.Second)
			if p == nil {
				fmt.Printf("B完成所有产品的运输！\n")
				return
			}
			productChan.Remove(p)
			lock.Unlock()
			for {
				bIndex += 1
				fmt.Printf("B将产品【%s】放入第【%d】个盒子！\n", p.Value, bIndex)
				if bIndex == 9 {
					fmt.Printf("B对产品【%s】完成整个流程的运输！\n", p.Value)
					bIndex = 0
					break
				}
				wait <- 1
			}
		}
	}()
	end.Wait()

}

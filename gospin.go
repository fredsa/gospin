package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func main() {
	var wgStart sync.WaitGroup
	var wgDone sync.WaitGroup

	cores := runtime.NumCPU()
	fmt.Printf("Found %d CPU cores.\n", cores)
	count := cores - 1
	wgStart.Add(count)
	wgDone.Add(count)
	for i := 0; i < count; i++ {
		go func(instance int) {
			fmt.Printf("Spinning up instance %d…\n", instance)
			defer wgDone.Done()
			wgStart.Done()
			for {
				if rand.Intn(100000000*count) < 1 {
					fmt.Print(".")
				}
			}
		}(i)
	}
	fmt.Printf("Spinning up %d cores…\n", count)
	wgStart.Wait()
	fmt.Printf("%d cores spinning :)\n", count)
	wgDone.Wait()
}

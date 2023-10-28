package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
)

func main() {
	var wgStart sync.WaitGroup
	var wgDone sync.WaitGroup

	args := os.Args

	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s <num_free_cores>\n\n", filepath.Base(args[0]))
		os.Exit(1)
	}

	freecores, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid number of core(s): %s\n", args[1])
		os.Exit(1)
	}

	cores := runtime.NumCPU()
	spincount := cores - freecores
	if spincount <= 0 {
		fmt.Fprintf(os.Stderr, "Requested free core(s) %d greater than available core(s) %d\n", freecores, cores)
		os.Exit(1)
	}

	fmt.Printf("Found %d CPU core(s). Will save %d core(s).\n", cores, freecores)
	wgStart.Add(spincount)
	wgDone.Add(spincount)
	for i := 0; i < spincount; i++ {
		go func(instance int) {
			fmt.Printf("Spinning up instance %d…\n", instance)
			defer wgDone.Done()
			wgStart.Done()
			for {
				if rand.Intn(100000000*spincount) < 1 {
					fmt.Print(".")
				}
			}
		}(i)
	}
	fmt.Printf("Spinning up %d go spin routines…\n", spincount)
	wgStart.Wait()
	fmt.Printf("%d go routine(s) spinning :)\n", spincount)
	wgDone.Wait()
}

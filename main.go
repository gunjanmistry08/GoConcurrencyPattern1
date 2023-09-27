package main

import (
	"fmt"
	"sync"
)

func NoArgs() {
	var wg sync.WaitGroup
	m := []string{"Hello", "World", "Happiness"}
	for _, v := range m {
		wg.Add(1)
		go func() {
			fmt.Println(v)
			defer wg.Done()
		}()
	}
	wg.Wait()
}
func Args() {
	var wg sync.WaitGroup
	m := []string{"Hello", "World", "Happiness"}
	for _, v := range m {
		wg.Add(1)
		go func(s string) {
			fmt.Println(s)
			defer wg.Done()
		}(v)
	}
	wg.Wait()
}

func main() {
	fmt.Println("Hello World")
	NoArgs()
	Args()
}

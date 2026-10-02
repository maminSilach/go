package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2)
	go func() {
		defer close(ch)
		time.Sleep(30 * time.Millisecond)
		ch <- "slow"
	}()

	go func() {
		defer close(ch)
		time.Sleep(10 * time.Millisecond)
		ch <- "fast"
	}()

	select {
	case v := <-ch:
		fmt.Print(v)
	case <-time.After(100 * time.Millisecond):

		fmt.Print("timeout")
	}
}

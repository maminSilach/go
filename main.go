package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		nums[i], _ = strconv.Atoi(sc.Text())
	}

	var mu sync.Mutex
	var wg sync.WaitGroup
	total := 0

	chunks := 4
	chunkSize := (n + chunks - 1) / chunks

	for i := 0; i < chunks; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > n {
			end = n
		}
		if start >= n {
			start = n
			end = n
		}

		wg.Add(1)
		go func(part []int) {
			defer wg.Done()

			sum := 0
			for _, v := range part {
				sum += v
			}

			mu.Lock()
			total += sum
			mu.Unlock()
		}(nums[start:end])
	}

	wg.Wait()
	fmt.Println(total)
}

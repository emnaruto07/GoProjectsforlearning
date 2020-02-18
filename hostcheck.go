package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {
	jobs := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()
			for domain := range jobs {
				_, err := net.LookupHost(domain)
				if err != nil {
					continue
				}

				fmt.Println(domain)
			}
		}()

	}

	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		domain := reader.Text()
		jobs <- domain

	}
	close(jobs)
	wg.Wait()
}

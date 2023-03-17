package main

import (
	"fmt"
	"sync"
)

func printInterface(coba []interface{}, bisa []interface{}, wg *sync.WaitGroup, mu *sync.Mutex, tipe string) {
	state := "interface1"
	wg.Add(8)
loop:
	for i := 1; i <= 4; i++ {
		if tipe == "rapih" {
			mu.Lock()
		}
		go func(i int) {
			defer wg.Done()
			if state == "interface1" {
				fmt.Println(bisa, i)
			} else {
				fmt.Println(coba, i)
			}
			if tipe == "rapih" {
				mu.Unlock()
			}
		}(i)
		if i == 4 && state == "interface1" {
			state = "interface2"
			goto loop
		}
	}
	wg.Wait()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	coba := []interface{}{"coba1", "coba2", "coba3"}
	bisa := []interface{}{"bisa1", "bisa2", "bisa3"}
	//Print Interface Acak
	printInterface(coba, bisa, &wg, &mu, "acak")
	fmt.Println()
	//Print Interface Rapih
	printInterface(coba, bisa, &wg, &mu, "rapih")
}

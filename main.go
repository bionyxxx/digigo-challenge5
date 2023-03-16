package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func processInterface1(data interface{}, wg *sync.WaitGroup, index int) {
	defer wg.Done()

	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Println(data, index)
}

func processInterface2(data interface{}, wg *sync.WaitGroup, index int) {
	defer wg.Done()

	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Println(data, index)
}

func main() {
	var wg sync.WaitGroup

	data1 := []string{"coba1", "coba2", "coba3"}
	data2 := []string{"bisa1", "bisa2", "bisa3"}

	for i := 0; i < 4; i++ {
		wg.Add(2)

		go func(data interface{}) {
			processInterface1(data, &wg, i)
			processInterface2(data, &wg, i)
			processInterface2(data, &wg, i)
			processInterface1(data, &wg, i)
		}(data1)

		go func(data interface{}) {
			processInterface2(data, &wg, i)
			processInterface1(data, &wg, i)
			processInterface2(data, &wg, i)
			processInterface1(data, &wg, i)
		}(data2)
	}

	wg.Wait()
}

/*package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mu sync.Mutex

func processData(data interface{}, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 4; i++ {
		mu.Lock()
		if rand.Intn(2) == 0 {
			fmt.Printf("%v %v\n", data, i)
			time.Sleep(100 * time.Millisecond)
		} else {
			fmt.Printf("%v %v\n", data, i)
			time.Sleep(150 * time.Millisecond)
		}
		mu.Unlock()
	}
}

func main() {
	data1 := []string{"bisa1", "bisa2", "bisa3"}
	data2 := []string{"coba1", "coba2", "coba3"}

	var wg sync.WaitGroup

	for i := 0; i < 1; i++ {
		wg.Add(2)
		go processData(data1, i+1, &wg)
		go processData(data2, i+1, &wg)
	}

	wg.Wait()
}*/

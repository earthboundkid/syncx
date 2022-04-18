package syncx_test

import (
	"fmt"
	"sync"

	"github.com/carlmjohnson/syncx"
)

func ExampleMutex() {
	n := syncx.NewMutex(2)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go n.Lock(func(n *int) {
			*n = *n * *n
			wg.Done()
		})
	}
	wg.Wait()
	fmt.Println(n.Read())
	// Output:
	// 4294967296
}

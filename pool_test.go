package syncx_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/carlmjohnson/syncx"
)

func ExamplePool() {
	// Make a bytes pool that gives out 100 byte chunks
	var p syncx.Pool[[]byte]
	p.New = func() *[]byte {
		b := make([]byte, 0, 100)
		return &b
	}
	p.Reset = func(p *[]byte) {
		*p = (*p)[:0]
	}

	// Use the chunks them copy
	// the bytes into some sink
	var buf strings.Builder
	f := func(i int) {
		b := p.Get()
		defer p.Put(b)

		*b = strconv.AppendInt(*b, int64(i), 10)
		*b = append(*b, ", "...)
		buf.Write(*b)
	}
	// Measure the allocs
	allocs := testing.AllocsPerRun(10, func() {
		for i := 0; i < 100; i++ {
			f(i + 1)
		}
	})
	// Look ma, no allocs!
	fmt.Println(allocs)
	// Output:
	// 0
}

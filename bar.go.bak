package pbar

import (
	"fmt"
	//"io"
	"sync"
)

type Bar struct {
	Counter    int
	CounterMax int
	Max64      int64
	Current64  int64
}

var WGpbar sync.WaitGroup
var pbar Bar

func NewBar(n int64) *Bar {
	WGpbar = sync.WaitGroup{}
	WGpbar.Add(1)

	return &Bar{
		Counter:   0,
		Max64:     n,
		Current64: 0,
	}

}

func (b *Bar) WithMax64(n int64) *Bar {
	b.Max64 = n
	return b
}

func (b *Bar) Add64(n int64) error {
	b.Current64 = b.Current64 + n
	b.Counter = b.Counter + 1

	if b.Counter < 10000 {
		return nil
	}

	if b.Counter%10000 == 0 {
		fmt.Printf("\r[Processing: %v / %v bytes | %v | %v ]", b.Current64, b.Max64, b.Counter, n)
		return nil
	}

	if b.Current64 >= b.Max64 {
		fmt.Printf("\r[Done: %v / %v bytes | %v | %v ]", b.Current64, b.Max64, b.Counter, n)
		fmt.Println("")
	}

	return nil
}

func (b *Bar) Write(bt []byte) (n int, err error) {
	n = len(bt)
	b.Add64(int64(n))
	return
}

func (b *Bar) Read(bt []byte) (n int, err error) {
	n = len(bt)
	b.Add64(int64(n))
	return
}

func (b *Bar) Finish() {
	WGpbar.Done()
}

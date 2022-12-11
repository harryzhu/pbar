package pbar

import (
	"fmt"

	//"io"
	"sync"
)

type Bar64 struct {
	Counter    int
	CounterMax int
	Max64      int64
	Current64  int64
}

var wg64 sync.WaitGroup
var bar64 *Bar64

func NewBar64(n int64) *Bar64 {
	wg64 = sync.WaitGroup{}
	wg64.Add(1)

	bar64 = &Bar64{
		Counter:   0,
		Max64:     n,
		Current64: 0,
	}

	return bar64
}

func (b *Bar64) WithMax64(n int64) *Bar64 {
	b.Max64 = n
	return b
}

func (b *Bar64) Add64(n int64) error {
	b.Current64 = b.Current64 + n
	b.Counter = b.Counter + 1

	if b.Counter < 10000 {
		return nil
	}

	if b.Counter%5000 == 0 {
		b.Render64("Processing")
		return nil
	}

	if b.Current64 >= b.Max64 {
		b.Render64("Done")
		fmt.Println("")

		wg64.Done()
	}

	return nil
}

func (b *Bar64) Render64(s string) error {
	if b.Max64 > 0 {
		fmt.Printf("\r[ %12s: %v / %v bytes ]", s, b.Current64, b.Max64)
	} else {
		fmt.Printf("\r[ %12s: %v bytes ]", s, b.Current64)
	}
	return nil
}

func (b *Bar64) Write(bt []byte) (n int, err error) {
	n = len(bt)
	b.Add64(int64(n))
	return
}

func (b *Bar64) Read(bt []byte) (n int, err error) {
	n = len(bt)
	b.Add64(int64(n))
	return
}

func (b *Bar64) Finish() {

}

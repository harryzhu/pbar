package pbar

import (
	"fmt"
	"time"

	"sync/atomic"
)

type Bar64 struct {
	Counter     int
	Max64       int64
	Current64   int64
	TimeStart64 int64
	TimeStop64  int64
	Speed64     int64
}

func NewBar64(n int64) *Bar64 {
	return &Bar64{
		Counter:     0,
		Max64:       n,
		Current64:   0,
		TimeStart64: 0,
		TimeStop64:  0,
		Speed64:     0,
	}
}

func (b *Bar64) WithMax64(n int64) *Bar64 {
	b.Max64 = n
	return b
}

func (b *Bar64) Add64(n int64) error {
	b.Current64 = atomic.AddInt64(&b.Current64, n)
	b.Counter = b.Counter + 1

	if b.Counter < 10000 {
		return nil
	}

	if b.Counter%5000 == 0 {
		b.TimeStop64 = time.Now().Unix()
		b.Render64("Processing")
		return nil
	}

	return nil
}

func (b *Bar64) Render64(s string) error {
	dSecond := b.TimeStop64 - b.TimeStart64
	if dSecond <= 0 {
		dSecond = 1
	}
	b.Speed64 = b.Current64 / dSecond / int64(1<<20)
	if b.Max64 > 0 {
		fmt.Printf("\r%12s:[ %v / %v bytes | speed: %v MB/s]", s, b.Current64, b.Max64, b.Speed64)
	} else {
		fmt.Printf("\r%12s:[ %v MB | speed: %v MB/s]", s, b.Current64/int64(1<<20), b.Speed64)
	}
	return nil
}

func (b *Bar64) Write(bt []byte) (n int, err error) {
	if b.Counter == 0 {
		b.TimeStart64 = time.Now().Unix()
	}
	n = len(bt)
	b.Add64(int64(n))
	return n, nil
}

func (b *Bar64) Read(bt []byte) (n int, err error) {
	if b.Counter == 0 {
		b.TimeStart64 = time.Now().Unix()
	}
	n = len(bt)
	b.Add64(int64(n))
	return n, nil
}

func (b *Bar64) Finish() error {
	b.Render64("Done")

	fmt.Println("")
	return nil
}

package pbar

import (
	"fmt"
	"time"

	"sync/atomic"
)

var (
	dSecond64 int64
)

type Bar64 struct {
	Disabled64   bool
	Counter      int
	CounterSkip  int
	CounterCycle int
	Max64        int64
	Current64    int64
	TimeStart64  int64
	TimeStop64   int64
	Speed64      int64
}

func NewBar64(n int64) *Bar64 {
	return &Bar64{
		Disabled64:   false,
		Counter:      0,
		CounterSkip:  1000,
		CounterCycle: 500,
		Max64:        n,
		Current64:    0,
		TimeStart64:  0,
		TimeStop64:   0,
		Speed64:      0,
	}
}

func (b *Bar64) WithMax64(n int64) *Bar64 {
	b.Max64 = n
	return b
}

func (b *Bar64) WithDisabled64(n bool) *Bar64 {
	b.Disabled64 = n
	return b
}

func (b *Bar64) WithCounterSkip(n int) *Bar64 {
	b.CounterSkip = n
	return b
}

func (b *Bar64) WithCounterCycle(n int) *Bar64 {
	b.CounterCycle = n
	return b
}

func (b *Bar64) Add64(n int64) error {
	if b.Disabled64 {
		return nil
	}
	b.Current64 = atomic.AddInt64(&b.Current64, n)
	b.Counter = b.Counter + 1

	if b.Counter < b.CounterSkip {
		return nil
	}

	if b.CounterCycle > 0 && b.Counter%b.CounterCycle == 0 {
		b.TimeStop64 = time.Now().Unix()
		b.Render64("Processing")
		return nil
	}

	return nil
}

func (b *Bar64) Render64(s string) error {
	if b.Disabled64 {
		return nil
	}
	dSecond64 = b.TimeStop64 - b.TimeStart64
	if dSecond64 <= 0 {
		dSecond64 = 1
	}
	b.Speed64 = b.Current64 / dSecond64 / int64(1<<20)
	if b.Max64 > 0 {
		fmt.Printf("\r%12s:[ %v / %v bytes | speed: %v MB/s]", s, b.Current64, b.Max64, b.Speed64)
	} else {
		fmt.Printf("\r%12s:[ %v MB | speed: %v MB/s]", s, b.Current64/int64(1<<20), b.Speed64)
	}
	return nil
}

func (b *Bar64) Write(bt []byte) (n int, err error) {
	if b.Disabled64 {
		return len(bt), nil
	}

	if b.Counter == 0 {
		b.TimeStart64 = time.Now().Unix()
	}

	n = len(bt)
	b.Add64(int64(n))
	return n, nil
}

func (b *Bar64) Read(bt []byte) (n int, err error) {
	if b.Disabled64 {
		return len(bt), nil
	}
	if b.Counter == 0 {
		b.TimeStart64 = time.Now().Unix()
	}
	n = len(bt)
	b.Add64(int64(n))
	return n, nil
}

func (b *Bar64) Finish() error {
	if b.Disabled64 {
		return nil
	}
	if b.Counter < b.CounterSkip {
		return nil
	}

	b.Render64("Done")
	fmt.Println("")
	return nil
}

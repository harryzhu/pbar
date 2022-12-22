package pbar

import (
	"fmt"
	"time"
)

var (
	dSecond int64
)

type Bar struct {
	Disabled     bool
	Counter      int
	CounterSkip  int
	CounterCycle int
	TimeStart    int64
	TimeStop     int64
	Max          int
	Current      int
	Speed        int64
}

func NewBar(n int) *Bar {
	return &Bar{
		Disabled:     false,
		Counter:      0,
		CounterSkip:  100,
		CounterCycle: 50,
		TimeStart:    0,
		TimeStop:     0,
		Max:          n,
		Current:      0,
		Speed:        0,
	}
}

func (b *Bar) WithMax(n int) *Bar {
	b.Max = n
	return b
}

func (b *Bar) WithDisabled(n bool) *Bar {
	b.Disabled = n
	return b
}

func (b *Bar) WithCounterSkip(n int) *Bar {
	b.CounterSkip = n
	return b
}

func (b *Bar) WithCounterCycle(n int) *Bar {
	b.CounterCycle = n
	return b
}

func (b *Bar) Add(n int) error {
	if b.Disabled {
		return nil
	}
	if b.Counter == 0 {
		b.TimeStart = time.Now().Unix()
	}
	b.Current = b.Current + n
	b.Counter = b.Counter + 1

	if b.Counter < b.CounterSkip {
		return nil
	}

	if b.CounterCycle > 0 && b.Counter%b.CounterCycle == 0 {
		b.TimeStop = time.Now().Unix()
		b.Render("Processing")
		return nil
	}

	return nil
}

func (b *Bar) Render(s string) error {
	if b.Disabled {
		return nil
	}

	dSecond = b.TimeStop - b.TimeStart
	if dSecond <= 0 {
		dSecond = 1
	}
	b.Speed = dSecond

	if b.Max > 0 {
		fmt.Printf("\r%12s:[ %v / %v | %v seconds]", s, b.Current, b.Max, b.Speed)
	} else {
		fmt.Printf("\r%12s:[ %v | %v seconds]", s, b.Current, b.Speed)
	}
	return nil
}

func (b *Bar) Finish() error {
	if b.Disabled {
		return nil
	}

	if b.Counter < b.CounterSkip {
		return nil
	}

	b.Render("Done")
	fmt.Println("")
	return nil
}

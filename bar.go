package pbar

import (
	"fmt"
)

type Bar struct {
	Counter      int
	CounterSkip  int
	CounterCycle int
	Max          int
	Current      int
}

func NewBar(n int) *Bar {
	return &Bar{
		Counter:      0,
		CounterSkip:  100,
		CounterCycle: 50,
		Max:          n,
		Current:      0,
	}
}

func (b *Bar) WithMax(n int) *Bar {
	b.Max = n
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
	b.Current = b.Current + n
	b.Counter = b.Counter + 1

	if b.Counter < b.CounterSkip {
		return nil
	}

	if b.Counter%b.CounterCycle == 0 {
		b.Render("Processing")
		return nil
	}

	return nil
}

func (b *Bar) Render(s string) error {
	if b.Max > 0 {
		fmt.Printf("\r%12s:[ %v / %v ]", s, b.Current, b.Max)
	} else {
		fmt.Printf("\r%12s:[ %v ]", s, b.Current)
	}
	return nil
}

func (b *Bar) Finish() error {
	if b.Counter < b.CounterSkip {
		return nil
	}

	b.Render("Done")
	fmt.Println("")
	return nil
}

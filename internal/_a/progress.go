package _a

import (
	"log"
	"sync/atomic"
	"time"
)

type Progress struct {
	count  *int64
	total  int64
	ticker *time.Ticker
	start  time.Time
	stop   chan struct{}
}

func NewProgress(interval time.Duration, total int64) *Progress {
	p := &Progress{
		ticker: time.NewTicker(interval),
		total:  total,
		start:  time.Now(),
		count:  new(int64),
		stop:   make(chan struct{}),
	}
	go p.print()
	return p
}

func (p *Progress) Add(c int64) {
	atomic.AddInt64(p.count, c)
}

func (p *Progress) Inc() {
	p.Add(1)
}

func (p *Progress) Get() int64 {
	return atomic.LoadInt64(p.count)
}

func (p *Progress) Stop() {
	p.ticker.Stop()
	close(p.stop)
	p.doPrint()
	log.Printf("Total duration: %v\n", time.Now().Sub(p.start))
}

func (p *Progress) print() {
	for {
		select {
		case <-p.ticker.C:
			p.doPrint()
		case <-p.stop:
			return
		}
	}
}

func (p *Progress) doPrint() {
	ct := p.Get()
	rate := int(float64(ct) / time.Now().Sub(p.start).Seconds())
	if p.total == 0 {
		log.Printf("Progress: %v; %v/s\n", ct, rate)
	} else {
		est := time.Second * time.Duration(int(p.total)/rate)
		log.Printf("Progress: %v/%v -- %.2f%%; %v/s; ETA %v\n", ct, p.total, float64(ct)/float64(p.total), rate, est)
	}
}

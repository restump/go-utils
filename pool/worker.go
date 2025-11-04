package poolutils

import (
	"sync"
)

type Work interface {
	Do()
	Error() error
}

type WorkerPool struct {
	workset chan Work
	waitgrp *sync.WaitGroup
}

func NewWorkerPool(size int) *WorkerPool {
	p := &WorkerPool{}
	p.workset = make(chan Work)
	p.waitgrp = &sync.WaitGroup{}

	for range size {
		p.waitgrp.Go(func() {
			for work := range p.workset {
				work.Do()
			}
		})
	}
	return p
}

func (p *WorkerPool) AddWork(work Work) {
	p.workset <- work
}

func (p *WorkerPool) Wait() {
	close(p.workset)
	p.waitgrp.Wait()
}

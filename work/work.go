package work

import "sync"

type Worker interface {
	Do()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	pool := Pool{
		work: make(chan Worker),
	}
	pool.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range pool.work {
				w.Do()
			}
			pool.wg.Done()
		}()
	}
	return &pool
}

func (p *Pool) Run(w Worker) {
	p.work <- w
}

func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}

package main

import (
	"github.com/alextuan1024/work/work"
	"log"
	"sync"
	"time"
)

func main() {
	pool := work.New(5)
	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			p := printer{
				name: name,
			}
			go func() {
				pool.Run(&p)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	pool.Shutdown()
}

var names = []string{
	"steve",
	"tony",
	"thor",
	"natasha",
	"clint",
	"bruce",
	"nick",
	"phil",
	"hill",
}

type printer struct {
	name string
}

func (m *printer) Do() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

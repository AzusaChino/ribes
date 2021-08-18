package ribes

import "time"

type Runnable interface {
	run()
}

// Pool from ThreadPoolExecutor
type Pool struct {
	corePoolSize    int
	maximumPoolSize int
	keepAliveTime   time.Duration
	workQueue       workingQueue
}

func (p *Pool) Execute(r Runnable) {

	if r == nil {
		panic("nil runnable executed")
	}

	if WorkerCount(p) < p.corePoolSize {

	}
}

func WorkerCount(p *Pool) int {
	return len(p.workQueue.works)
}

func (p *Pool) addWorker(runnable Runnable, core bool) bool {
retry:
	for {
		p.workQueue.add(runnable)
		if o := recover(); o != nil {
			goto retry
		}
	}
}

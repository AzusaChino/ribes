package pool

import (
    "github.com/azusachino/ribes/lock"
    "sync"
    "sync/atomic"
)

type Runnable interface {
    run()
}

// Pool from ThreadPoolExecutor
type Pool struct {
    corePoolSize int // max workers
    lock         *sync.Locker
    cond         *sync.Cond
    workQueue    workingQueue
    status       uint64
}

func NewPool(coreSize int) *Pool {
    lk := lock.NewSpinLock()
    wq := workingQueue{}
    st := new(uint64)
    atomic.StoreUint64(st, 1)
    return &Pool{
        corePoolSize: coreSize,
        lock:         &lk,
        workQueue:    wq,
        status:       *st,
    }
}

func (p *Pool) Execute(r Runnable) {

    if r == nil {
        panic("nil runnable executed")
    }

    if WorkerCount(p) < p.corePoolSize {
        r.run()
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

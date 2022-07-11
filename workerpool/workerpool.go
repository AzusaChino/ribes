package workerpool

import (
    "log"
    "runtime"
    "sync"
    "time"
)

type Runnable func(...interface{})

type workerPool struct {
    WorkerFunc Runnable
    MaxWorkersCount int
    MaxIdleWorkerDuration time.Duration

    lock sync.Mutex
    workersCount int
    mustStop bool

    ready []*workerChan
    stopCh chan struct{}

    workerChanPool sync.Pool
}

type workerChan struct {
    lastUseTime time.Time
    ch chan interface{}
}

var workerChanCap = func() int {
    if runtime.GOMAXPROCS(0) == 1 {
        return 0
    }
    return 1
}()

func (wp * workerPool) Start() {
    if wp.stopCh != nil {
        log.Panic("")
    }
    wp.stopCh = make(chan struct{})

    stopCh := wp.stopCh
    wp.workerChanPool.New = func() interface{} {
        return &workerChan{ch: make(chan interface{}, workerChanCap)}
    }

    go func() {
        //var scratch []*workerChan
        for {
            select {
            case <-stopCh:
                return
            default:
                time.Sleep(wp.MaxIdleWorkerDuration)
            }
        }
    }()
}
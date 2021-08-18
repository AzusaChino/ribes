package ribes

type workingQueue struct {
	works []Runnable
}

func (wq *workingQueue) add(r Runnable) {
	wq.works = append(wq.works, r)
}

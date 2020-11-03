package pool

import "log"

// WorkerChannel allows communication between dispatcher and workers
var WorkerChannel = make(chan chan Work)

// Collector collects jobs and distributes them to workers
type Collector struct {
	Work chan Work // receives jobs and sends to workers
	End  chan bool // to stop workers
}

// StartDispatcher instantiates and connects all of the workers with the worker pool
func StartDispatcher(workerCount int) Collector {
	input := make(chan Work) // channel to receive work
	end := make(chan bool)

	collector := Collector{
		Work: input,
		End:  end,
	}

	var i int
	var workers []Worker

	for i < workerCount {
		i++
		log.Println("starting worker:", i)
		worker := Worker{
			ID:            i,
			JobChannel:    make(chan Work),
			WorkerChannel: WorkerChannel,
			End:           make(chan bool),
		}
		worker.Start()
		workers = append(workers, worker) // stores worker
	}

	// start collector
	go func() {
		for {
			select {
			case <-end:
				for _, w := range workers {
					w.Stop() // stop worker
				}
				return
			case work := <-input:
				worker := <-WorkerChannel // wait for available channel
				worker <- work            // dispatch work to worker
			}
		}
	}()

	return collector
}

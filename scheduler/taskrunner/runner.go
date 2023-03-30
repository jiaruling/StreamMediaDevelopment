package taskrunner

type Runner struct {
	Controller controlChan
	Error      controlChan
	Data       dataChan
	dataSize   int
	longlived  bool
	Dispatcher fn
	Exector    fn
}

func NewRunner(size int, longlived bool, d, e fn) *Runner {
	return &Runner{
		Controller: make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		dataSize:   size,
		longlived:  longlived,
		Dispatcher: d,
		Exector:    e,
	}
}

func (r *Runner) startDispatcher() {
	defer func() {
		if !r.longlived {
			close(r.Controller)
			close(r.Error)
			close(r.Data)
		}
	}()

	for {
		select {
		case c := <-r.Controller:
			if c == READY_TO_DISPATCH {
				if err := r.Dispatcher(r.Data); err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_EXECUTE
				}
			}
			if c == READY_TO_EXECUTE {
				if err := r.Exector(r.Data); err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}
		case e := <-r.Error:
			if e == CLOSE {
				return
			}
		}
	}
}

func (r *Runner) StartAll() {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatcher()
}

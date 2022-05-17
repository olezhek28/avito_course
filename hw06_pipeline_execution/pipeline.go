package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if in == nil {
		return nil
	}

	for _, st := range stages {
		tmpChan := make(chan interface{})
		go func(ch In) {
			defer close(tmpChan)

			for {
				select {
				case item, ok := <-ch:
					if !ok {
						return
					}
					tmpChan <- item
				case <-done:
					return
				}
			}
		}(in)

		in = st(tmpChan)
	}

	return in
}

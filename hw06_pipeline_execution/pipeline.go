package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, st := range stages {
		tmpChan := make(chan interface{})
		go func(ch In) {
			defer close(tmpChan)
			for data := range ch {
				tmpChan <- data
			}
		}(in)

		in = st(tmpChan)
	}

	return in
}

//func runStage(ctx context.Context, in In, stage Stage) Out {
//	out := stage(in)
//
//	go func() {
//		//for data := range in {
//		//
//		//}
//		select {
//		case data := <-in:
//			out <- stage(data)
//		case <-ctx.Done():
//			return
//		}
//	}()
//
//	return out
//}

package hw06_pipeline_execution // nolint:golint,stylecheck

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := make(Bi)
	stageInOut := in

	for _, stage := range stages {
		stageInOut = stage(stageInOut)
	}

	go func() {
		defer close(out)

		for {
			select {
			case v, ok := <-stageInOut:
				if !ok {
					return
				}

				out <- v
			case <-done:
				return
			}
		}
	}()

	return out
}

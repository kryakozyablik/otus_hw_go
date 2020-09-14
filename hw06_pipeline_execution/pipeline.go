package hw06_pipeline_execution // nolint:golint,stylecheck

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		stageIn := make(Bi)

		go func(in In, stageIn Bi) {
			defer close(stageIn)

			for {
				select {
				case v, ok := <-in:
					if !ok {
						return
					}

					stageIn <- v
				case <-done:
					return
				}
			}
		}(in, stageIn)

		in = stage(stageIn)
	}

	return in
}

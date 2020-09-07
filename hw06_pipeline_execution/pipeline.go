package hw06_pipeline_execution // nolint:golint,stylecheck

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := make(Bi)
	inBi := make(Bi)

	go func() {
		defer close(inBi)

		for {
			select {
			case v, ok := <-in:
				if !ok {
					return
				}

				inBi <- v
			case <-done:
				return
			}
		}
	}()

	result := exec(inBi, stages...)

	go func() {
		defer close(out)

		for {
			select {
			case v, ok := <-result:
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

func exec(in In, stages ...Stage) Out {
	stageInOut := in

	for _, stage := range stages {
		stageInOut = stage(stageInOut)
	}

	return stageInOut
}

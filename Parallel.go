package ipo

import (
	"fmt"
	"sync"
)

// ParallelOutputs tries to write each output parallelly and returns when all outputs finished.
func ParallelOutputs(obj interface{}, outputs []Output) error {
	var wg sync.WaitGroup

	for _, output := range outputs {
		wg.Add(1)

		go func(out Output) {
			defer wg.Done()

			err := out.Write(obj)

			if err != nil {
				fmt.Println(err)
			}
		}(output)
	}

	wg.Wait()
	return nil
}

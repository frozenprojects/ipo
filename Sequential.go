package ipo

import (
	"fmt"
	"sync"
)

// SequentialInputs tries each input sequentially and returns the first successful one.
func SequentialInputs(inputs []Input) (obj interface{}, err error) {
	for _, input := range inputs {
		obj, err = input.Read()

		if err == nil {
			return obj, nil
		}
	}

	return nil, err
}

// SequentialOutputs tries to write each output sequentially and returns when one fails.
func SequentialOutputs(obj interface{}, outputs []Output) error {
	for _, output := range outputs {
		err := output.Write(obj)

		if err != nil {
			return err
		}
	}

	return nil
}

// ParallelOutputs tries to write each output parallelly and returns when all outputs finished.
func ParallelOutputs(obj interface{}, outputs []Output) error {
	var wg sync.WaitGroup

	for _, output := range outputs {
		wg.Add(1)

		go func(out Output) {
			err := out.Write(obj)

			if err != nil {
				fmt.Println(err)
			}

			wg.Done()
		}(output)
	}

	wg.Wait()
	return nil
}

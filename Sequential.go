package ipo

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

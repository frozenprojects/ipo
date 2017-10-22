package ipo

// System represents a full IPO system.
type System struct {
	Inputs          []Input
	Outputs         []Output
	InputProcessor  InputProcessor
	OutputProcessor OutputProcessor
}

// Run ...
func (system *System) Run(data interface{}) error {
	obj, err := system.InputProcessor.Process(system.Inputs)

	if err != nil {
		return err
	}

	return system.OutputProcessor.Process(obj, system.Outputs)
}

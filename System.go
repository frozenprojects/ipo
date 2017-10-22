package ipo

// System represents a full IPO system.
type System struct {
	Inputs          []Input
	Outputs         []Output
	InputProcessor  func(inputs []Input) (obj interface{}, err error)
	OutputProcessor func(obj interface{}, outputs []Output) error
}

// Run runs the system once for the given data.
func (system *System) Run() error {
	obj, err := system.InputProcessor(system.Inputs)

	if err != nil {
		return err
	}

	return system.OutputProcessor(obj, system.Outputs)
}

package ipo

// Processor processes all inputs and returns a single object which is fed to the outputs.
type Processor interface {
	Process([]Input) (interface{}, error)
}

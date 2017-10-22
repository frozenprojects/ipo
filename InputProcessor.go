package ipo

// InputProcessor processes all inputs and returns a single object which is fed to the outputs.
type InputProcessor interface {
	Process(inputs []Input) (obj interface{}, err error)
}

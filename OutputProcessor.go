package ipo

// OutputProcessor implements the logic to handle the outputs.
type OutputProcessor interface {
	Process(obj interface{}, outputs []Output) error
}

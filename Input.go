package ipo

// Input represents a system that can fetch an object somewhere (e.g. from the network).
type Input interface {
	Read() (obj interface{}, err error)
}

package ipo

// Output represents a system that saves the object somewhere (e.g. to the filesystem).
type Output interface {
	Write(interface{}) error
}

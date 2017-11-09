package inputs

import "image"

// Image ...
type Image interface {
	// Image ...
	Image() image.Image

	// Data ...
	Data() []byte

	// Format ...
	Format() string

	// Extension ...
	Extension() string
}

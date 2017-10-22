package outputs

import (
	"fmt"
	"image"

	"github.com/aerogo/ipo"
)

// Force interface implementation
var _ ipo.Output = (*ImageFile)(nil)

// ImageFile writes an image to the filesystem.
type ImageFile struct {
	Directory string
}

// Write ...
func (file *ImageFile) Write(obj interface{}) error {
	img := obj.(image.Image)
	fmt.Println(file.Directory, img.Bounds().Dx(), img.Bounds().Dy())
	return nil
}

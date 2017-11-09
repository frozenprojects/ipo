package outputs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/aerogo/ipo/inputs"

	"github.com/aerogo/ipo"
)

// Force interface implementation
var _ ipo.Output = (*ImageFile)(nil)

// ImageFile writes an image to the filesystem.
type ImageFile struct {
	Directory string
	BaseName  string
	Format    string
	Size      int
}

// Write ...
func (file *ImageFile) Write(obj interface{}) error {
	networkImage := obj.(*inputs.NetworkImage)

	extension := networkImage.Extension()

	if extension == "" {
		return errors.New("Unknown format: " + networkImage.Format())
	}

	img := networkImage.Image()
	fmt.Println(file.BaseName+extension, "|", img.Bounds().Dx(), "x", img.Bounds().Dy(), "|", len(networkImage.Data())/1024, "KB")

	// Original file output
	if file.Format == "" {
		fullPath := path.Join(file.Directory, file.BaseName+extension)
		return ioutil.WriteFile(fullPath, networkImage.Data(), 0644)
	}

	return nil
}

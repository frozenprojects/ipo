package outputs

import (
	"errors"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"path"

	"github.com/aerogo/ipo/inputs"
	"github.com/chai2010/webp"

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
	Quality   float32
}

// Write ...
func (file *ImageFile) Write(obj interface{}) error {
	networkImage := obj.(*inputs.NetworkImage)

	extension := networkImage.Extension()

	if extension == "" {
		return errors.New("Unknown format: " + networkImage.Format())
	}

	img := networkImage.Image()
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	resizeRequired := file.Size != 0 && file.Size != width

	fmt.Println(file.BaseName+extension, "|", width, "x", height, "|", len(networkImage.Data())/1024, "KB")

	// Original file output
	if file.Format == "" && !resizeRequired {
		fullPath := path.Join(file.Directory, file.BaseName+extension)
		return ioutil.WriteFile(fullPath, networkImage.Data(), 0644)
	}

	if file.Format == "webp" {
		fullPath := path.Join(file.Directory, file.BaseName+".webp")
		return SaveWebP(img, fullPath, file.Quality)
	}

	return nil
}

// SaveWebP saves an image as a file in WebP format.
func SaveWebP(img image.Image, out string, quality float32) error {
	file, writeErr := os.Create(out)

	if writeErr != nil {
		return writeErr
	}

	defer file.Close()

	encodeErr := webp.Encode(file, img, &webp.Options{
		Quality: quality,
	})

	return encodeErr
}

package outputs

import (
	"errors"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path"

	"github.com/aerogo/ipo/inputs"
	"github.com/chai2010/webp"
	"github.com/nfnt/resize"

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
	Quality   int
}

// Write ...
func (file *ImageFile) Write(obj interface{}) error {
	input := obj.(inputs.Image)
	extension := input.Extension()

	if extension == "" {
		return errors.New("Unknown format: " + input.Format())
	}

	img := input.Image()
	width := img.Bounds().Dx()
	resizeRequired := file.Size != 0 && file.Size < width

	// fmt.Println(file.BaseName+extension, "|", width, "x", height, "|", len(input.Data())/1024, "KB")

	// Original file output
	if file.Format == "" && !resizeRequired {
		fullPath := path.Join(file.Directory, file.BaseName+extension)
		return ioutil.WriteFile(fullPath, input.Data(), 0644)
	}

	// Resize if needed
	if resizeRequired {
		img = resize.Resize(uint(file.Size), 0, img, resize.Lanczos3)
	}

	// Set format automatically if needed
	if file.Format == "" {
		file.Format = input.Format()
	}

	// Write data to file
	fullPath := path.Join(file.Directory, file.BaseName+file.Extension())
	stream, err := os.Create(fullPath)

	if err != nil {
		return err
	}

	defer stream.Close()

	switch file.Format {
	case "jpg", "jpeg":
		err = jpeg.Encode(stream, img, &jpeg.Options{
			Quality: file.Quality,
		})
	case "png":
		err = png.Encode(stream, img)
	case "gif":
		err = gif.Encode(stream, img, nil)
	case "webp":
		err = webp.Encode(stream, img, &webp.Options{
			Quality: float32(file.Quality),
		})
	}

	return err
}

// Extension ...
func (file *ImageFile) Extension() string {
	switch file.Format {
	case "jpg", "jpeg":
		return ".jpg"
	case "png":
		return ".png"
	case "gif":
		return ".gif"
	case "webp":
		return ".webp"
	default:
		return ""
	}
}

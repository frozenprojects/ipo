package ipo_test

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aerogo/ipo"
	"github.com/aerogo/ipo/inputs"
	"github.com/aerogo/ipo/outputs"
)

func TestImageDownload(t *testing.T) {
	system := &ipo.System{
		Inputs: []ipo.Input{
			&inputs.NetworkImage{
				URL: "https://blitzprog.org/images/icons/countries/gb.png",
			},
		},
		Outputs: []ipo.Output{
			&outputs.ImageFile{
				Directory: "./",
			},
		},
		InputProcessor:  ipo.SequentialInputs,
		OutputProcessor: ipo.SequentialOutputs,
	}

	err := system.Run()

	assert.NoError(t, err)
}

package canvas

import (
	"github.com/stretchr/testify/require"
	"ray-tracer-challenge/geometry"
	"testing"
)

func TestCanvas_NewCanvas(t *testing.T) {
	canvas := NewCanvas(10, 20)

	require.NotNil(t, canvas)
	require.Equal(t, 10, canvas.Width)
	require.Equal(t, 20, canvas.Height)
	for i := 0; i < canvas.Height; i++ {
		for j := 0; j < canvas.Width; j++ {
			require.Equal(t, geometry.Tuple{
				0.0,
				0.0,
				0.0,
				geometry.ColorW}, canvas.Pixels[i][j])
		}
	}
}

func TestCanvas_WritePixel(t *testing.T) {
	canvas := NewCanvas(10, 20)
	color := geometry.NewColor(1, 0, 0)

	canvas.WritePixel(2, 3, color)
	sut := canvas.PixelAt(2, 3)
	require.Equal(t, color, sut)
}

func TestCanvas_ConstructPPM(t *testing.T) {
	canvas := NewCanvas(5, 3)
	sut := canvas.ToPPM()

	require.Contains(t, sut, "P3\n5 3\n255\n")
}

func TestCanvas_PPMColors(t *testing.T) {
	canvas := NewCanvas(5, 3)
	canvas.WritePixel(0, 0, geometry.NewColor(1.5, 0, 0))
	canvas.WritePixel(1, 2, geometry.NewColor(0, 0.5, 0))
	canvas.WritePixel(2, 4, geometry.NewColor(-0.5, 0, 1))

	expected := "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0\n" +
		"0 0 0 0 0 0 0 128 0 0 0 0 0 0 0\n" +
		"0 0 0 0 0 0 0 0 0 0 0 0 0 0 255\n"

	require.Contains(t, canvas.ToPPM(), expected)
}

// TODO cleanup ppm generation
/*
func TestCanvas_PPMColorsLineLength(t *testing.T) {
	canvas := NewCanvas(10, 2, geometry.NewColor(1, 0.8, 0.6))

	expected := "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204\n" +
		"153 255 204 153 255 204 153 255 204 153 255 204 153\n" +
		"255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204\n" +
		"153 255 204 153 255 204 153 255 204 153 255 204 153\n"

	require.Contains(t, canvas.ToPPM(), expected)
}
*/
